package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitMQ interface {
	Connection() *amqp.Connection
	Channel() *amqp.Channel
}

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ() IRabbitMQ {
	rabbitmq := &RabbitMQ{}

	rabbitmq.setConnection()
	rabbitmq.setQueue()

	return rabbitmq
}
func (r *RabbitMQ) Connection() *amqp.Connection {
	return r.conn
}

func (r *RabbitMQ) Channel() *amqp.Channel {
	return r.setChannel()
}

func (r *RabbitMQ) setConnection() {
	conn, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		log.Fatal("Error to execute rabbitMQ")
	}
	r.conn = conn
}

func (r *RabbitMQ) setChannel() (c *amqp.Channel) {
	c, err := r.conn.Channel()
	if err != nil {
		log.Fatal("Error to execute rabbitMQ")
	}
	return
}

func (r *RabbitMQ) setQueue() {
	q, _ := r.Channel().QueueDeclare(
		os.Getenv("AMQP_QUEUE"), // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	r.queue = q
}
