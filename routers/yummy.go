package routers

import (
	"net/http"
	"onaka-api/cruds"

	"github.com/gin-gonic/gin"
)

func initYummyRouter(yr *gin.RouterGroup) {
	yr.Use(middleware)
	yr.POST("/:post_id", lookLikeYummy)
	yr.DELETE("/:post_id", doNotLookLikeYummy)
}

func lookLikeYummy(c *gin.Context) {
	var (
		userId  any
		isExist bool
	)

	postId := c.Param("post_id")

	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	post, err := cruds.GiveYummy(postId, userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func doNotLookLikeYummy(c *gin.Context) {
	var (
		userId  any
		isExist bool
	)

	postId := c.Param("post_id")

	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	post, err := cruds.DeleteYummy(postId, userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}
