package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
	utils "github.com/usefss/cab-please/identity/utils"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var demandCacheWindowSize int = 60 * 1000

func PublishDecreaseDemand(demandKey string) {
	conn, err := amqp.Dial("amqp://guest:guest@" + utils.GetEnv("RABBITMQ_ADDR") + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"
	err = ch.ExchangeDeclare(
		utils.GetEnv("RABBITMQ_DELAYED_EXCHANGE_NAME"), // name
		"x-delayed-message",                            // type
		true,                                           // durable
		false,                                          // auto-deleted
		false,                                          // internal
		false,                                          // no-wait
		args,                                           // arguments
	)

	err = ch.Publish(
		utils.GetEnv("RABBITMQ_DELAYED_EXCHANGE_NAME"), // exchange
		utils.GetEnv("RABBITMQ_DECREASE_QUEUE_NAME"),   // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(demandKey),
			Headers: amqp.Table{
				"x-delay": demandCacheWindowSize,
			},
		})
	failOnError(err, "Failed to publish a message")
}
