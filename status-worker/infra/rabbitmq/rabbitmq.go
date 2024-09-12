package rabbitmq

import (
	"os"
	"time"

	"github.com/Arturlima/status-worker/utils"
	"github.com/rabbitmq/amqp091-go"
)

type IRabbitMQ interface {
	Connection() *amqp091.Connection
	Channel() *amqp091.Channel
	Reconnect()
}

type RabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   amqp091.Queue
}

func NewRabbitMQ() IRabbitMQ {
	rabbitmq := &RabbitMQ{}

	rabbitmq.setConnection()
	rabbitmq.setQueue()

	return rabbitmq

}

func (r *RabbitMQ) Connection() *amqp091.Connection {
	return r.conn
}

func (r *RabbitMQ) Channel() *amqp091.Channel {
	return r.setChannel()
}

func (r *RabbitMQ) setChannel() (c *amqp091.Channel) {
	c, err := r.conn.Channel()
	if err != nil {
		utils.FailWithError("failure creating channe", err)
		return
	}

	c.Confirm(false)

	err = c.Qos(1, 0, false)
	if err != nil {
		utils.FailWithError("failure to set Qos", err)
		return
	}

	return
}

func (r *RabbitMQ) Reconnect() {
	r.setConnection()
	r.setQueue()
}

func (r *RabbitMQ) setConnection() {
	amqpConfig := amqp091.Config{
		Vhost:     "/",
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
	}

	c, err := amqp091.DialConfig(os.Getenv("AMQP_SERVER_URL"), amqpConfig)
	if err != nil {
		utils.FailWithError("failure connecting to RabbitMQ", err)
	}

	r.conn = c
}

func (r *RabbitMQ) setQueue() {
	q, _ := r.Channel().QueueDeclare(
		os.Getenv("AMQP_QUEUE"),
		true,
		false,
		false,
		false,
		nil)

	r.queue = q
}
