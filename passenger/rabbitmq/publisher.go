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
	conn, errDial := amqp.Dial("amqp://guest:guest@" + utils.GetEnv("RABBITMQ_ADDR") + "/")
	failOnError(errDial, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, errChann := conn.Channel()
	failOnError(errChann, "Failed to open a channel")
	defer ch.Close()

	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"
	errEx := ch.ExchangeDeclare(
		utils.GetEnv("RABBITMQ_DELAYED_EXCHANGE_NAME"), // name
		"x-delayed-message",                            // type
		true,                                           // durable
		false,                                          // auto-deleted
		false,                                          // internal
		false,                                          // no-wait
		args,                                           // arguments
	)
	failOnError(errEx, "Failed to declare an exchange")

	errPub := ch.Publish(
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
	failOnError(errPub, "Failed to publish a message")
}
