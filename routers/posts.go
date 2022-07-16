package routers

import (
	"net/http"
	"onaka-api/cruds"
	"onaka-api/db"
	"onaka-api/types"

	"github.com/gin-gonic/gin"
)

func initPostRouter(pr *gin.RouterGroup) {
	pr.Use(middleware)
	pr.GET("", getPosts)
	pr.GET("/:post_id", getPostById)
	pr.POST("", postPosts)
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

// func postPostsImage(c *gin.Context){
// 	file, err := c.FormFile("file")
// 	cruds.PostImages(file)
// 	form, _ := c.MultipartForm()
//     files := form.File["file"]
//     configPath := filepath.Join(".", "volley", "csv")
// 	//file, fileHeader, err := c.Request.FormFile("file")
// }

func postPosts(c *gin.Context) {
	var (
		userId  any
		isExist bool
		p       db.Post
		err     error
	)

	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}
	var payload types.CreatePost
	c.Bind(&payload)
	if p, err = cruds.PostPosts(payload.Content, payload.ImageUrl, userId.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &p)
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
