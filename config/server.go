package config

func GetAMQPURI() string {
	//this must be passed in or created by the app based on the bindings
	return "amqp://guest:guest@localhost:5672/"
}
