package logic

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/zeromicro/go-zero/core/logx"
	"stock/internal/model"
	"stock/internal/svc"
	"stock/service/order"
	"strconv"
	"time"
)

// ConsumeDeductMessage 消费扣减库存的消息
// TODO: 保证消费消息的过程不允许出错，不然消息会被提交而导致数据不一致，可能需要类似死信队列，然后让人工处理。
func ConsumeDeductMessage(svcCtx *svc.ServiceContext) {
	err := svcCtx.Consumer.Subscribe("stock-deduct", nil)
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
				res, err := svcCtx.Consumer.Commit()
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
		msg, err := svcCtx.Consumer.ReadMessage(-1)
		if err != nil {
			logx.Errorf("Consumer ReadMessage Err:%v, msg(%v)\n", err, msg)
			continue // NOTE:
		}

		var orderId int
		err = json.Unmarshal(msg.Value, &orderId)
		if err != nil {
			logx.Error("queue consume err: ", err)
			continue // NOTE: 需要保证Unmarshal不出错
		}

		// 更新stock表
		res, err := svcCtx.OrderRpc.GetOrder(context.Background(), &order.OrderModel{Id: int64(orderId)})
		if err == nil { // 意味着这个消息需要去更新库存
			res, err := svcCtx.StockModel.DecrOne(context.Background(), int(res.StockId))
			if err != nil {
				logx.Error("queue consume err: ", err) // NOTE: 不允许出错
				continue
			} else {
				affected, err := res.RowsAffected()
				if err != nil {
					logx.Error("queue consume err: ", err)
					continue
				}
				if affected != 1 {
					logx.Error("queue consume err: affected != 1")
				}
			}
		} else {
			logx.Error(err)
		}

		Cnt += 1
		if Cnt >= batchCommitSize {
			Cnt = 0
			ch <- struct{}{}
		}
	}
}

func SendDeductMessage(client *kafka.Producer, task *model.StockTask, topic string) error {
	data, err := json.Marshal(task)
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
		Key:   []byte(strconv.Itoa(int(task.Id))),
		Value: data,
	}

	// TODO: 这里是瓶颈。为了提升速度和并发量，这里采用了批提交，宕机将会丢失还在本地缓冲区的数据
	err = client.Produce(msg, nil)
	if err != nil {
		return err
	}

	return nil
}
