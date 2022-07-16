package routers

import (
	"net/http"
	"onaka-api/cruds"

	"github.com/gin-gonic/gin"
)

func initFunnyRouter(fr *gin.RouterGroup) {
	fr.Use(middleware)
	fr.POST("/:post_id", itIsFunny)
	fr.DELETE("/:post_id", itIsNotFunny)
}

func itIsFunny(c *gin.Context) {
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

	post, err := cruds.GiveFunny(postId, userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func itIsNotFunny(c *gin.Context) {
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

	post, err := cruds.DeleteFunny(postId, userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}
