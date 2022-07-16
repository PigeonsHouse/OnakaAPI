package routers

import (
	"net/http"
	"onaka-api/cruds"
	"onaka-api/types"

	"github.com/gin-gonic/gin"
)

func initUserRouter(ur *gin.RouterGroup) {
	ur.POST("/signup", signUp)
	ur.POST("/signin", signIn)
}

func signUp(c *gin.Context) {
	var payload types.SignUpUser
	c.Bind(&payload)

	u, err := cruds.CreateUser(payload.Name, payload.Email, payload.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email is already exist",
		})
		return
	}

	c.JSON(http.StatusOK, &u)
}

func signIn(c *gin.Context) {
	var payload types.SignInUser
	c.Bind(&payload)

	u, err := cruds.GenerateJWT(payload.Email, payload.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &u)
}
