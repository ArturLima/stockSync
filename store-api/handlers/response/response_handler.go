package response

import (
	"net/http"

	"github.com/Arturlima/store-api/models"
	"github.com/gin-gonic/gin"
)

func SendOk(g *gin.Context, obj interface{}) {
	g.JSON(http.StatusOK, obj)
}

func SendError(g *gin.Context, statusCode int, err error) {
	g.Set(models.ExceptionError, true)
	g.Status(statusCode)
	g.Error(err)
	g.Abort()
}
