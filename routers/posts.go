package routers

import (
	"net/http"
	"onaka-api/cruds"

	"github.com/gin-gonic/gin"
)

func initPostRouter(pr *gin.RouterGroup) {
	pr.Use(middleware)
	pr.GET("", getPosts)
	pr.GET("/:post_id", getPostById)
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

func getPostById(c *gin.Context) {
	postId := c.Param("post_id")
	post, err := cruds.GetPost(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}
