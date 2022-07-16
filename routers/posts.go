package routers

import (
	"net/http"
	"onaka-api/cruds"

	"github.com/gin-gonic/gin"
)

func initPostRouter(pr *gin.RouterGroup) {
	pr.GET("", getPosts)
}

func getPosts(c *gin.Context) {
	timeline, err := cruds.GetTimeLine()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, timeline)
}
