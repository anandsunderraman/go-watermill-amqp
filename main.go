// Sources for https://watermill.io/docs/getting-started/
package main

import (
	"context"
	"go-watermill-amqp/config"
)

func main() {



	//this must be passed in or created by the app based on the bindings
	var amqpURI = "amqp://guest:guest@localhost:5672/"

	amqpSubscriber, err := config.GetAMQPSubscriber(amqpURI)

	if err != nil {
		panic(err)
	}

	router, err := config.GetRouter()

	if err != nil {
		panic(err)
	}

	config.ConfigureAMQPSubscriptionHandlers(router, amqpSubscriber)

	ctx := context.Background()
	if err = router.Run(ctx); err != nil {
		panic(err)
	}
}




