package config

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"go-watermill-amqp/handlers"
)

func GetRouter() (*message.Router, error) {

	logger := watermill.NewStdLogger(false, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}
	return router, nil
}

func ConfigureAMQPSubscriptionHandlers(r *message.Router, s message.Subscriber) {

	r.AddNoPublisherHandler(
		"OnLightMeasured", // handler name, must be unique
		"light/measured",  // topic from which we will read events
		s,
		handlers.OnLightMeasured, // topic to which we will publish events
	)
}
