package broker

// To avoid hash conflict the Broker enum follows the amount of the database enum,
// so the max value from one is subtract from the other
const (
	RabbitMQ = "6"
	Kafka    = "7"
)
