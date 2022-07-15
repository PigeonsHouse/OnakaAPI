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

	v1 := api.Group("/api/v1")
	user_router := v1.Group("/users")
	initUserRouter(user_router)
	post_router := v1.Group("/posts")
	initPostRouter(post_router)
	yummy_router := v1.Group("/yummy")
	initYummyRouter(yummy_router)
	funny_router := v1.Group("/funny")
	initFunnyRouter(funny_router)
}
