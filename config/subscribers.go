package config

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

func GetAMQPSubscriber(amqpURI string) (*amqp.Subscriber, error) {
	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)

	amqpSubscriber, err := amqp.NewSubscriber(
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}
	return amqpSubscriber, nil
}
