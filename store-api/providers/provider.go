package providers

import (
	"github.com/Arturlima/store-api/handlers"
	"github.com/Arturlima/store-api/handlers/controllers"
	"github.com/Arturlima/store-api/infra/rabbitmq"
)

type IProvider interface {
	ScopedStoreController() *controllers.StoreController
}

type Provider struct{}

func NewProvider() IProvider {
	return &Provider{}
}

func (p *Provider) ScopedStoreController() *controllers.StoreController {
	pub := rabbitmq.NewPublisher()
	handler := handlers.NewStoreHandler(pub)

	return controllers.NewStoreController(handler)
}
