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

func PublishDecreaseDemand(demandKey string) {
	conn, err := amqp.Dial("amqp://guest:guest@" + utils.GetEnv("RABBITMQ_ADDR") + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		utils.GetEnv("DecreaseDemandQueueName"), // name
		false,                                   // durable
		false,                                   // delete when unused
		false,                                   // exclusive
		false,                                   // no-wait
		nil,                                     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(demandKey),
		})
	failOnError(err, "Failed to publish a message")
}
