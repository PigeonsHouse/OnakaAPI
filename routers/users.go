package routers

import (
	"fmt"
	"net/http"
	"onaka-api/cruds"
	"onaka-api/db"
	"onaka-api/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initUserRouter(ur *gin.RouterGroup) {
	ur.POST("/signup", signUp)
	ur.POST("/signin", signIn)
	ur.GET("/@me", middleware, getMe)
	ur.PATCH("/@me", middleware, updateName)
	ur.DELETE("/@me", middleware, deleteMe)
	ur.GET("/:user_id", getUser)
	ur.GET("/@me/posts", middleware, getMyPosts)
	ur.GET("/:user_id/posts", middleware, getSomeonesPosts)
}

func signUp(c *gin.Context) {
	var payload types.SignUpUser
	c.Bind(&payload)

	u, err := cruds.CreateUser(payload.Name, payload.Email, payload.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &u)
}

func signIn(c *gin.Context) {
	var payload types.SignInUser
	c.Bind(&payload)
	fmt.Printf("%s %s\n", payload.Email, payload.Password)

	u, err := cruds.GenerateJWT(payload.Email, payload.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &u)
}

func getMe(c *gin.Context) {
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

	userInfo := &db.User{}
	if err := cruds.GetUserByID(userInfo, userId.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user is not exist",
		})
		return
	}

	c.JSON(http.StatusOK, userInfo)
	return
}

func getUser(c *gin.Context) {
	userId := c.Param("user_id")
	userInfo := &db.User{}
	if err := cruds.GetUserByID(userInfo, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user is not exist",
		})
		return
	}

	c.JSON(http.StatusOK, userInfo)
	return
}

func deleteMe(c *gin.Context) {
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

	if err := cruds.DeleteUser(userId.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func updateName(c *gin.Context) {
	var (
		userId  any
		isExist bool
		user    db.User
		err     error
	)

	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	name := c.Query("name")
	if user, err = cruds.UpdateName(userId.(string), name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func getSomeonesPosts(c *gin.Context) {
	var (
		userId  any
		isExist bool
		posts   []db.Post
		err     error
	)

	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit <= 0 {
		limit = 50
	}
	if page <= 0 {
		page = 1
	}

	userId = c.Param("user_id")
	if posts, err = cruds.GetPostsByUserId(userId.(string), limit, page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func getMyPosts(c *gin.Context) {
	var (
		userId  any
		isExist bool
		posts   []db.Post
		err     error
	)

	if userId, isExist = c.Get("user_id"); !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit <= 0 {
		limit = 50
	}
	if page <= 0 {
		page = 1
	}

	if posts, err = cruds.GetPostsByUserId(userId.(string), limit, page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
