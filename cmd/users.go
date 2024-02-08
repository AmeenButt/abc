package cmd

import (
	"github.com/ashbeelghouri/project1/middlewares"
	"github.com/ashbeelghouri/project1/repo"
	"github.com/gin-gonic/gin"
)

func CreateUserRoutes(router *gin.Engine) {
	group := router.Group("/users")
	{
		group.POST("/create-users", middlewares.CreateUser, repo.CreateUser)
	}
}
