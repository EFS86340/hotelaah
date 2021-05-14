package hotelaah

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// the specific DS for this project used in func Publish()
type AahData interface {

	// should be transferred to kafka message
	Value() []byte

}

type KafkeeCallBack interface {

}

// A kafka producer has 3 mandatory properties:
// 1. bootstrap.servers (better more than one
// 2. key.serializer
// 3. value.serializer

// There are 3 primary methods of sending messages
// 1. Fire-and-forget
// 2. Synchronous send
// 3. Async send

// the producer to write data to kafka
type Kafkee struct {
	broker string
	topic string

	ok bool 
	pd *kafka.Producer
	// TODO: whether channel should be declared and defined else where
	deliveryChan chan kafka.Event
}

func NewKafkee(t, b string) *Kafkee {
	return &Kafkee{
		broker: b,
		topic: t,
		ok: false,
	}
}

func (k *Kafkee) Init() error {
	var err error
	// we should not use shorthand :=, thus shadows the in-struct var
	k.pd, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": k.broker})
	if err != nil {
		log.Printf("[kafka] open Producer to %s failed", k.broker)
	} else {
		k.ok = true
		log.Printf("[kafka] open Producer to %s success.", k.broker)
	}
	return err
}

func (k *Kafkee) IsOk() bool {
	return k.ok
}

// publish a message to topic
func (k *Kafkee) Publish(data AahData) {
	err := k.pd.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &(k.topic), Partition: kafka.PartitionAny},
	Value: data.Value(),
	Headers: []kafka.Header{{Key: "testHeader", Value: []byte("Headers value")}},
	}, k.deliveryChan)
				 if err != nil {

				 }

}

func (k *Kafkee) PublishWithConfirm(data AahData) {

}

func (k *Kafkee) PublishBatch(datas []AahData) {

}

// Disconnect from kafka broker
func (k *Kafkee) Disconnect() bool {
	return true
}


type KafkeeManager struct {
	kees []Kafkee
}
