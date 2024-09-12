package worker

import (
	"log"

	"github.com/Arturlima/status-worker/core/handlers"
	"github.com/Arturlima/status-worker/infra/rabbitmq"
)

type IWorker interface {
	StartWorker()
}

type Worker struct {
	handler  handlers.IPackageHandler
	consumer rabbitmq.IConsumer
}

func NewWorker() IWorker {
	return &Worker{
		handler:  handlers.NewPackageHandler(),
		consumer: rabbitmq.NewConsumer(),
	}
}

func (w *Worker) StartWorker() {

	messages, err := w.consumer.Delivery()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sucessfully connected to RabbitMQ")
	log.Println("Waiting for messages...")

	var forever chan struct{}
	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)
			if done := w.handler.AddOrUpdateStatus(message.Body); done {
				message.Ack(false)
			}
		}
	}()
	<-forever

}
