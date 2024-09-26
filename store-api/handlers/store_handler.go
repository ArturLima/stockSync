package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Arturlima/store-api/handlers/requests"
	"github.com/Arturlima/store-api/handlers/response"
	"github.com/Arturlima/store-api/infra/rabbitmq"
	"github.com/Arturlima/store-api/pkg/utils"
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

	err := h.pub.Publish(body)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err)
		return
	}

	response.SendOk(c, utils.AsResult(body))
}
