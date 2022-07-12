package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}

func InitRouter(api *gin.Engine) {
	api.GET("/", hello)
}
