package routers

import (
	"fmt"
	"net/http"
	"strings"

	"onaka-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}

func InitRouter(api *gin.Engine) {
	api.GET("/", hello)

	v1 := api.Group("/api/v1")
	user_router := v1.Group("/users")
	initUserRouter(user_router)
	post_router := v1.Group("/posts")
	initPostRouter(post_router)
	yummy_router := v1.Group("/yummy")
	initYummyRouter(yummy_router)
	funny_router := v1.Group("/funny")
	initFunnyRouter(funny_router)
}

func middleware(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		ary := strings.Split(authorizationHeader, " ")
		if len(ary) == 2 {
			if ary[0] == "Bearer" {
				t, err := jwt.ParseWithClaims(ary[1], &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
					return utils.SigningKey, nil
				})

				if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
					userId := (*claims)["sub"].(string)
					c.Set("user_id", userId)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
	c.Next()
}
