package service

import (
	"encoding/json"
	"os"

	"github.com/emersonluiz/go-user/logger"
	"github.com/emersonluiz/go-user/models"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.SetLog(msg + " :: " + err.Error())
	}
}

func SendMessage(user *models.User) {
	conn, err := amqp.Dial(os.Getenv("RABBIT_HOST"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("RABBIT_QUEUE"), // name
		false,                     // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	jsonUser, err := json.Marshal(user)
	if err != nil {
		logger.SetLog(err.Error())
		return
	}

	body := jsonUser
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
