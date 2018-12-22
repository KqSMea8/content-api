package routes

import (
	"github.com/sundogrd/content-api/handler"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Gql(r *gin.Engine) {
	r.GET("/hello", api.Hello)
	r.GET("/graphql", api.gql)
}
