package rabbitmq

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/Arturlima/store-api/handlers/requests"
	"github.com/rabbitmq/amqp091-go"
)

type IPublisher interface {
	Publish(s requests.Product) (err error)
}

type Publisher struct {
	rabbitmq IRabbitMQ
}

func NewPublisher() IPublisher {
	return &Publisher{
		rabbitmq: NewRabbitMQ(),
	}
}

func (p *Publisher) Publish(s requests.Product) (err error) {
	m, _ := json.Marshal(s)
	message := amqp091.Publishing{
		ContentType: "application/json",
		Body:        []byte(m),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.rabbitmq.Channel().PublishWithContext(
		ctx,
		"",
		os.Getenv("AMQP_QUEUE"),
		false,
		false,
		message,
	)
	return
}
