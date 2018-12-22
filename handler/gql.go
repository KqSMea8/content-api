package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	gqlgenHandler "github.com/99designs/gqlgen/handler"
)

// Hello ...
func Gql(c *gin.Context) {
	http.Handle("/query", gqlgenHandler.GraphQL(main.NewExecutableSchema(main.Config{Resolvers: &main.Resolver{}})))
	c.JSON(200, gin.H{
		"msg": "hello world!",
	})
}