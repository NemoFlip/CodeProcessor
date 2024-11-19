package workers

import (
	"HomeWork1/configs"
	"HomeWork1/internal/entity"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func SendCode(data entity.CodeRequest) {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Printf("unable to parse config file: %s", err)
		return
	}
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Printf("Can't run the RabbitMQ server: %s", err.Error())
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open channel: %s", err.Error())
		return
	}
	defer func() {
		_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	}()
	queue, err := ch.QueueDeclare(
		"codeProcessor",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %s", err.Error())
		return
	}
	codeJSON, _ := json.Marshal(data)
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        codeJSON,
		},
	)
	if err != nil {
		log.Printf("Unable to send a message: %s", err.Error())
		return
	}

}
