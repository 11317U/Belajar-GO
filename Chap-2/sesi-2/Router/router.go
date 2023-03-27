package router

import (
	controllers "sesi-2/Controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	route := gin.Default()

	route.POST("/Book", controllers.CreateBook)
	route.PUT("/Book/:bookID", controllers.UpadateBook)
	route.GET("/Book/:bookID", controllers.GetBook)
	route.GET("/Book", controllers.GetAllBook)
	route.DELETE("/Book/:bookID", controllers.DeleteBook)

	return route

}
