package controllers

import (
	"github.com/gin-gonic/gin"
)

//Login login
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	c.JSON(200, gin.H{
		"status":   "posted",
		"username": username,
		"password": password,
	})
}
