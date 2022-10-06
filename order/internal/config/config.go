package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string `json:"DataSource"`
	} `json:"DB"`

	KafkaCommon struct {
		ApiVersionRequest bool   `json:"ApiVersionRequest"`
		BootstrapServers  string `json:"BootstrapServers"`
		SecurityProtocol  string `json:"SecurityProtocol"`
		SslCaLocation     string `json:"SslCaLocation"`
		SaslMechanism     string `json:"SaslMechanism"`
		SaslUsername      string `json:"SaslUsername"`
		SaslPassword      string `json:"SaslPassword"`
	} `json:"KafkaCommon"`

	KafkaProducer struct {
		Acks                       string `json:"Acks"`
		BatchSize                  int    `json:"BatchSize"`
		CompressionType            string `json:"CompressionType"`
		DeliveryTimeoutMs          int    `json:"DeliveryTimeoutMs"`
		LingerMs                   int    `json:"LingerMs"`
		MessageMaxBytes            int    `json:"MessageMaxBytes"`
		Retries                    int    `json:"Retries"`
		RetryBackoffMs             int    `json:"RetryBackoffMs"`
		StickyPartitioningLingerMs int    `json:"StickyPartitioningLingerMs"`
	} `json:"KafkaProducer"`

	KafkaConsumer struct {
		AutoOffsetReset        string `json:"AutoOffsetReset"`
		EnableAutoCommit       string `json:"EnableAutoCommit"`
		FetchMaxBytes          int    `json:"FetchMaxBytes"`
		GroupId                string `json:"GroupId"`
		HeartbeatIntervalMs    int    `json:"HeartbeatIntervalMs"`
		MaxPollIntervalMs      int    `json:"MaxPollIntervalMs"`
		MaxPartitionFetchBytes int    `json:"MaxPartitionFetchBytes"`
		SessionTimeoutMs       int    `json:"SessionTimeoutMs"`
	} `json:"KafkaConsumer"`
}
