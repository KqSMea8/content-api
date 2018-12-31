package routes

import (
	"code.byted.org/learning_fe/pathfinder-api/handler"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Hello(r *gin.Engine) {
	r.GET("/hello", handler.Hello)
}
