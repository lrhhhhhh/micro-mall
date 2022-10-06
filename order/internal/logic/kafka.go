package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/zeromicro/go-zero/core/logx"
	"order/internal/model"
	"order/internal/svc"
	"strconv"
	"time"
)

type Job struct {
	Id       int    `json:"id"`       // Job唯一标识ID，确保唯一
	Topic    string `json:"topic"`    // 真正投递的消费队列
	Body     string `json:"body"`     // Job消息体
	Delay    int64  `json:"delay"`    // Job需要延迟的时间, 单位：秒
	ExecTime int64  `json:"execTime"` // Job执行的时间, 单位：秒
}

func (j *Job) Validate() error {
	if j.Topic == "" || j.Delay < 0 {
		return fmt.Errorf("invalid job %+v", j)
	}
	if j.ExecTime < time.Now().Unix() {
		return errors.New("job already expire before send to queue")
	}
	return nil
}

// Consume 消费延迟关闭订单的消息
func Consume(svcCtx *svc.ServiceContext) {
	logx.Infof("subscribe topic: `order-cancel`")
	consumer := svcCtx.Consumer
	err := consumer.Subscribe("order-cancel", nil)
	if err != nil {
		panic(err)
	}

	ticker := time.Tick(time.Second)
	ch := make(chan struct{})
	// commit async
	go func() {
		for {
			select {
			case <-ch:
			case <-ticker:
				res, err := consumer.Commit()
				if err != nil {
					if err.Error() != "Local: No offset stored" {
						logx.Error(err)
					}
				} else {
					_ = res
					//fmt.Printf("consumer commit: %+v\n", res)
				}
			}
		}
	}()

	Cnt := 0
	batchCommitSize := 1000

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			logx.Errorf("Consumer ReadMessage Err:%v, msg(%v)\n", err, msg)
			continue // NOTE:
		}

		var job Job
		err = json.Unmarshal(msg.Value, &job)
		if err != nil {
			logx.Error(err)
			continue // NOTE: 需要保证Unmarshal不出错
		}

		logx.Infof("order cancel message: %+v", job)

		// 超时取消订单
		// TODO: 增加数据库和redis的库存
		err = svcCtx.OrderModel.Update(context.Background(), &model.Order{
			Id:        int64(job.Id),
			Status:    -2,
			UpdatedAt: time.Now().Unix(),
		})
		if err != nil {
			logx.Error(err, job)
		}

		Cnt += 1
		if Cnt >= batchCommitSize {
			Cnt = 0
			ch <- struct{}{}
		}
	}
}

func CancelOrderAfter15m(client *kafka.Producer, orderId, delay int, topic, body string) error {
	job := Job{
		Topic:    topic,
		Id:       orderId,
		Delay:    int64(delay),
		ExecTime: int64(delay) + time.Now().Unix(),
		Body:     body,
	}

	err := job.Validate()
	if err != nil {
		return err
	}

	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	delayTopic := "delay-15m"
	msg := &kafka.Message{
		Timestamp: time.Unix(job.ExecTime, 0),
		TopicPartition: kafka.TopicPartition{
			Topic:     &delayTopic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(strconv.Itoa(orderId)),
		Value: data,
	}

	//ch := make(chan kafka.Event)
	//err = client.Produce(msg, ch)
	//if err != nil {
	//	return err
	//}
	//
	////ev := <-ch
	////resp := ev.(*kafka.Message)
	////fmt.Println("message recv: ", resp)
	//
	//<-ch

	// TODO: 这里是瓶颈。为了提升速度和并发量，这里采用了批提交，宕机将会丢失还在本地缓冲区的数据
	err = client.Produce(msg, nil)
	if err != nil {
		return err
	}

	return nil
}

func SendDeductMessage(client *kafka.Producer, orderId int, topic string) error {
	data, err := json.Marshal(orderId)
	if err != nil {
		return err
	}

	copyTopic := topic
	msg := &kafka.Message{
		Timestamp: time.Now(),
		TopicPartition: kafka.TopicPartition{
			Topic:     &copyTopic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(strconv.Itoa(orderId)),
		Value: data,
	}

	// TODO: 这里是瓶颈。为了提升速度和并发量，这里采用了批提交，宕机将会丢失还在本地缓冲区的数据
	err = client.Produce(msg, nil)
	if err != nil {
		return err
	}

	return nil
}
