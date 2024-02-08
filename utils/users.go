package utils

import (
	models "github.com/ashbeelghouri/project1/model"
	"github.com/gin-gonic/gin"
)

func headerBinding(c *gin.Context, data any) bool {
	if err := c.ShouldBindHeader(&data); err != nil {
		return false
	}
	return true
}

func ApiKeyProvided(c *gin.Context) bool {
	var data models.ValidApi
	return headerBinding(c, data)
}

func UserIsNotLoggedIn(c *gin.Context) bool {
	var data models.UserLoggedIn
	return headerBinding(c, data)
}

func UserShouldbeLoggedIn(c *gin.Context) bool {
	var data models.UserLoggedIn
	return headerBinding(c, data)
}
