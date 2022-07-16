package routers

import (
	"fmt"
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
	fmt.Println(payload)

	u := cruds.CreateUser(payload.Name, payload.Email, payload.Password)

	c.JSON(http.StatusOK, &u)
}
