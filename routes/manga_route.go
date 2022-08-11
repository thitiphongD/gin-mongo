package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/gin-mongo/controllers"
)

func UserRoute(router *gin.Engine) {
	router.GET("/manga", controllers.GetAllManga())
	router.GET("/manga/:mangaID", controllers.GetManga())
	router.POST("/manga", controllers.CreateManga())
	router.PUT("/manga/:mangaID", controllers.EditManga())
	router.DELETE("/manga/:mangaID", controllers.DeleteManga())
}
