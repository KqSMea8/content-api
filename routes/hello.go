package routes

import (
	"github.com/sundogrd/creator-service/api"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Hello(r *gin.Engine) {
	r.GET("/hello", api.Hello)
	r.GET("/graphql", api.gql)
}
