package route

import (
	"github.com/gin-gonic/gin"
)

//Route Route
var Route *gin.Engine

func init() {
	Route = gin.Default()

	Route.GET("/", Login)
}
