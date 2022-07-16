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
	pr.DELETE("/:post_id", deletePost)
}

func getPosts(c *gin.Context) {
	if _, isExist := c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}
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
	if _, isExist := c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}
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

func deletePost(c *gin.Context) {
	var (
		userId  any
		isExist bool
	)
	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	postId := c.Param("post_id")
	err := cruds.DeletePost(postId, userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
