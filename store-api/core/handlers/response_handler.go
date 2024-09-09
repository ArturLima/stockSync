package handlers

import (
	"net/http"

	"github.com/Arturlima/store-api/core/enums"
	"github.com/gin-gonic/gin"
)

func sendOk(g *gin.Context, obj interface{}) {
	g.JSON(http.StatusOK, obj)
}

func sendError(g *gin.Context, statusCode int, err error) {
	g.Set(enums.ExceptionError, true)
	g.Status(statusCode)
	g.Error(err)
	g.Abort()
}
