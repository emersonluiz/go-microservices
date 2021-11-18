package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/emersonluiz/go-email/logger"
	"github.com/emersonluiz/go-email/models"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.SetLog(msg + " :: " + err.Error())
	}
}

func ReceiveMessage() {
	conn, err := amqp.Dial(os.Getenv("RABBIT_HOST"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("RABBIT_QUEUE"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var user *models.User
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &user)
			if err != nil {
				log.Fatal("Error on parser message")
			}
			SendEmail(user)
		}
	}()

	logger.SetLog(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
