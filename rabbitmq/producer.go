package rabbitmq

import (
	"HomeWork1/entity"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendCode(data entity.CodeRequest) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/") // Создаем подключение к RabbitMQ
	if err != nil {
		fmt.Printf("Can't run the RabbitMQ server: %s", err.Error())
		return
	}
	defer func() {
		_ = conn.Close() // Закрываем соединение в случае удачной попытки
	}()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to open channel: %s", err.Error())
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
		fmt.Printf("Failed to declare a queue: %s", err.Error())
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
		fmt.Printf("Unable to send a message: %s", err.Error())
		return
	}

}
