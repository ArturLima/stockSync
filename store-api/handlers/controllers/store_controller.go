package controllers

import (
	"github.com/Arturlima/store-api/handlers"
	"github.com/gin-gonic/gin"
)

type StoreController struct {
	handler handlers.IStoreHandler
}

func NewStoreController(h handlers.IStoreHandler) *StoreController {
	return &StoreController{
		handler: h,
	}
}

func (c *StoreController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/v1/store")

	routes.POST("/request", c.handler.RequestProductFromCD)
}
