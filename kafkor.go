package hotelaah

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkorCallBack interface {

}

type Kafkor struct {
	broker string
	group string
	topics []string

	ok bool
	cs *kafka.Consumer

}

func NewKafkor(b, g string, ts []string) *Kafkor {
	return &Kafkor{
		broker: b,
		group: g,
		topics: ts,
	}
}

func (k *Kafkor) Init() error {
	var err error
	k.cs, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": k.broker,
			"broker.address.family": "v4",
			"group.id": k.group,
			"session.timeout.ms": 6000,
			"auto.offset.reset": "earliest",
				})
	if err != nil {
		log.Printf("[kafka] open Consumer to %s failed.", k.broker)
	} else {
		k.ok = true;
		// TODO: if subscribe failure
		err = k.cs.SubscribeTopics(k.topics, nil)
		log.Printf("[kafka] open Consumer to %s success.", k.broker)
	}
	return err
}

func (k *Kafkor) IsOk() bool {
	return k.ok
}

func (k *Kafkor) Listen() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	for k.IsOk() {
		select {
			case sig := <-sigchan:
				log.Printf("[kafka] Kafkor received %v signal", sig)
				k.ok = false
			default:
			ev := k.cs.Poll(100)
						if ev == nil {
							continue
						}

					switch e := ev.(type) {
						case *kafka.Message:
						log.Printf("[kafka] receive Message on %s: %s", e.TopicPartition, string(e.Value))
							if e.Headers != nil {
								log.Printf("[kafka] Message Headers: %v\n", e.Headers)
							}
						case kafka.Error:
						log.Printf("[kafka] receive Error %v: %v\n", e.Code(), e)
							if e.Code() == kafka.ErrAllBrokersDown {
								k.ok = false
							}
						default:
						log.Printf("[kafka] Ignored: %v\n", e)
					}
		}
	}
	defer k.cs.Close()
}
