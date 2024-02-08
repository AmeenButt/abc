package middlewares

import (
	"net/http"

	"github.com/ashbeelghouri/project1/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	if !utils.ApiKeyProvided(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "false",
			"error":  "api unauthorized",
		})
	}

	if !utils.UserIsNotLoggedIn(c) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  "user must be logged in",
		})
	}

	if !utils.UserIsNotLoggedIn(c) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  "user must not be logged in",
		})
	}

	c.Next()

}
