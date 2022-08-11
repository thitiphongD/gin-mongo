package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/gin-mongo/configs"
	"github.com/thitiphongD/gin-mongo/routes"
)

func main() {
	router := gin.Default()

	// run database
	configs.CONNECT_DB()

	//routes
	routes.UserRoute(router)

	router.Run(":5000")
}
