package main

import (
	"fmt"
	"onaka-api/db"
	"onaka-api/routers"
	"onaka-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	utils.LoadEnv()
	db.InitDB()
	routers.InitRouter(api)
	api.Run(fmt.Sprintf(":%s", utils.ApiPort))
}
