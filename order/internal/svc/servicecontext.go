package svc

import (
	"database/sql"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"order/internal/config"
	"order/internal/model"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	Mysql      *sql.DB
	Consumer   *kafka.Consumer
	Producer   *kafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	db, err := conn.RawDB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1024)
	db.SetMaxIdleConns(20)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"auto.offset.reset":         c.KafkaConsumer.AutoOffsetReset,
		"enable.auto.commit":        c.KafkaConsumer.EnableAutoCommit,
		"fetch.max.bytes":           c.KafkaConsumer.FetchMaxBytes,
		"group.id":                  c.KafkaConsumer.GroupId,
		"heartbeat.interval.ms":     c.KafkaConsumer.HeartbeatIntervalMs,
		"max.partition.fetch.bytes": c.KafkaConsumer.MaxPartitionFetchBytes,
		"max.poll.interval.ms":      c.KafkaConsumer.MaxPollIntervalMs,
		"session.timeout.ms":        c.KafkaConsumer.SessionTimeoutMs,

		"api.version.request": c.KafkaCommon.ApiVersionRequest,
		"bootstrap.servers":   c.KafkaCommon.BootstrapServers,
		"security.protocol":   c.KafkaCommon.SecurityProtocol,
	})
	if err != nil {
		panic(err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"acks":                          c.KafkaProducer.Acks,
		"batch.size":                    c.KafkaProducer.BatchSize,
		"compression.type":              c.KafkaProducer.CompressionType,
		"delivery.timeout.ms":           c.KafkaProducer.DeliveryTimeoutMs,
		"linger.ms":                     c.KafkaProducer.LingerMs,
		"message.max.bytes":             c.KafkaProducer.MessageMaxBytes,
		"retries":                       c.KafkaProducer.Retries,
		"retry.backoff.ms":              c.KafkaProducer.RetryBackoffMs,
		"sticky.partitioning.linger.ms": c.KafkaProducer.StickyPartitioningLingerMs,

		"api.version.request": c.KafkaCommon.ApiVersionRequest,
		"bootstrap.servers":   c.KafkaCommon.BootstrapServers,
		"security.protocol":   c.KafkaCommon.SecurityProtocol,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		Mysql:      db,
		OrderModel: model.NewOrderModel(conn),
		Producer:   producer,
		Consumer:   consumer,
	}
}
