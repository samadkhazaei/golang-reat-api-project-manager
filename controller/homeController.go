package controller

import "github.com/gin-gonic/gin"

func HomeContent(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
