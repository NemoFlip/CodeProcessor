package workers

import (
	"HomeWork1/code_service/internal"
	"HomeWork1/configs"
	"HomeWork1/internal/entity"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func ConsumeMessage() []byte {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Printf("unable to parse config file: %s", err)
		return nil
	}
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Printf("Can't run the RabbitMQ server: %s", err.Error())
		return nil
	}
	defer func() {
		_ = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open channel: %s", err.Error())
		return nil
	}
	defer func() {
		_ = ch.Close()
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
		return nil
	}

	msg, ok, err := ch.Get(queue.Name, true)
	if err != nil {
		log.Printf("Failed to get a message: %s", err.Error())
	}
	if !ok {
		log.Printf("Consume task: was no delivery waiting or an error occured")
	}
	var codeInfo entity.CodeRequest
	_ = json.Unmarshal(msg.Body, &codeInfo)
	// Then we run user's code
	return app.RunCode(codeInfo)
}
