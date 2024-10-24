package app

import (
	"HomeWork1/configs"
	"HomeWork1/internal/entity"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessage() []byte {
	cfg, err := configs.GetConfig()
	if err != nil {
		fmt.Printf("unable to parse config file: %s", err)
		return nil
	}
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		fmt.Printf("Can't run the RabbitMQ server: %s", err.Error())
		return nil
	}
	defer func() {
		_ = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to open channel: %s", err.Error())
		return nil
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
		fmt.Printf("Failed to declare a queue: %s", err.Error())
		return nil
	}

	msg, ok, err := ch.Get(queue.Name, true)
	if err != nil {
		fmt.Printf("Failed to get a message: %s", err.Error())
	}
	if !ok {
		fmt.Printf("NOT OK!")
	}
	var codeInfo entity.CodeRequest
	_ = json.Unmarshal(msg.Body, &codeInfo)
	// Здесь запускаем пользовательский код
	return RunCode(codeInfo)
}
