package controller

import (
	"github.com/gin-gonic/gin"
	"samadkhazaei.dev/GinBuilerPlate/models"
)

func FindClients(c *gin.Context) {
	var clients []models.Client

	models.DB.Find(&clients)

	c.JSON(200, gin.H{"data": clients})
}
