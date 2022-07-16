package routers

import (
	"fmt"
	"net/http"
	"onaka-api/cruds"
	"onaka-api/db"
	"onaka-api/types"

	"github.com/gin-gonic/gin"
)

func initUserRouter(ur *gin.RouterGroup) {
	ur.POST("/signup", signUp)
	ur.POST("/signin", signIn)

	ur.GET("/@me", middleware, getMe)
	ur.DELETE("/@me", middleware, deleteMe)
	ur.GET("/:user_id", getUser)
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

	fmt.Println(userInfo)

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

	fmt.Println(userInfo)

	c.JSON(http.StatusOK, types.UserResponse{
		ID:        userInfo.ID,
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		CreatedAt: userInfo.CreatedAt,
		UpdatedAt: userInfo.UpdatedAt,
	})
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
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
