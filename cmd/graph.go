package cmd

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ashbeelghouri/project1/graph"
	"github.com/gin-gonic/gin"
)

func graphQLHandler() gin.HandlerFunc {
	log.Printf("inside graphql handler")
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CreateGraphqlRoutes(router *gin.Engine) {
	group := router.Group("/")
	{
		group.GET("/", playgroundHandler())
		group.POST("/query", graphQLHandler())
	}
}
