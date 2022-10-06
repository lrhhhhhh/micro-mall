package svc

import (
	"database/sql"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"stock/internal/config"
	"stock/internal/model"
	"stock/service/order"
	"stock/service/order/orderclient"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	StockModel     model.StockModel
	StockTaskModel model.StockTaskModel
	Mysql          *sql.DB
	Redis          *redis.Client
	Consumer       *kafka.Consumer
	Producer       *kafka.Producer
	OrderRpc       order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	db, err := conn.RawDB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4096)
	db.SetMaxIdleConns(2)

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

	opts := redis.Options{
		Addr:         c.Redis.Host,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolSize:     16,
		MinIdleConns: 4,
		//MaxConnAge:         0,
		//PoolTimeout:        0,
		//IdleTimeout:        0,
	}

	return &ServiceContext{
		Config:         c,
		StockModel:     model.NewStockModel(conn),
		StockTaskModel: model.NewStockTaskModel(conn),
		Mysql:          db,
		Redis:          redis.NewClient(&opts),
		Producer:       producer,
		Consumer:       consumer,
		OrderRpc:       orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
