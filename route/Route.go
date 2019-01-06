package route

import (
	"GoModDemo/controllers"

	"github.com/gin-gonic/gin"
)

//Route Route
var Route *gin.Engine

func init() {
	Route = gin.Default()

	Route.GET("/", controllers.Login)
}
