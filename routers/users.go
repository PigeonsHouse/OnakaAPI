package routers

import (
	"net/http"
	"onaka-api/cruds"
	"onaka-api/types"

	"github.com/gin-gonic/gin"
)

func initUserRouter(ur *gin.RouterGroup) {
	ur.POST("/signup", signUp)
}

func signUp(c *gin.Context) {
	var payload types.SignUpUser
	c.Bind(&payload)

	var u *types.UserResponse
	cruds.CreateUser(u, payload.Name, payload.Email, payload.Password)

	c.JSON(http.StatusOK, &u)
}
