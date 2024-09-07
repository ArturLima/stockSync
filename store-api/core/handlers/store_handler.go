package handlers

import (
	"log/slog"

	"github.com/Arturlima/store-api/core/handlers/requests"
	"github.com/Arturlima/store-api/infra/rabbitmq"
	"github.com/gin-gonic/gin"
)

type IStoreHandler interface {
	RequestProductFromCD(c *gin.Context)
}

type StoreHandler struct {
	pub rabbitmq.IPublisher
}

func NewStoreHandler(p rabbitmq.IPublisher) IStoreHandler {
	return &StoreHandler{
		pub: p,
	}
}

func (h *StoreHandler) RequestProductFromCD(c *gin.Context) {
	body := requests.Product{}

	if err := c.BindJSON(&body); err != nil {
		slog.Error("error bindJson", "error", err)
		return
	}

	return
}
