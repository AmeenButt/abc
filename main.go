package main

import (
	"context"

	"github.com/ashbeelghouri/project1/cmd"
	"github.com/ashbeelghouri/project1/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	router := gin.Default()

	conn := utils.ConnectDB(ctx)
	defer conn.Close(ctx)
	// cmd.CreateUserRoutes(router)

	cmd.CreateGraphqlRoutes(router)

	router.Run(":8080")
}
