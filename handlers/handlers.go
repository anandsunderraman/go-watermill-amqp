package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"go-watermill-amqp/payloads"
	"log"
)

func OnLightMeasuredX(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		var lm payloads.LightMeasured
		err := json.Unmarshal(msg.Payload, &lm)
		if err != nil {
			fmt.Printf("error unmarshalling message: %s, err is: %s", msg.Payload, err)
			msg.Nack()
		}

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func OnLightMeasured(msg *message.Message) error {

	log.Printf("received message payload: %s", string(msg.Payload))

	var lm payloads.LightMeasured
	err := json.Unmarshal(msg.Payload, &lm)
	if err != nil {
		fmt.Printf("error unmarshalling message: %s, err is: %s", msg.Payload, err)
		return err
	}

	return nil
}
