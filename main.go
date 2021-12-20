package main

import (
	"github.com/gin-gonic/gin"
	"samadkhazaei.dev/GinBuilerPlate/controller"
	"samadkhazaei.dev/GinBuilerPlate/models"
)

func main() {
	app := gin.New()

	models.ConnectDatabase()

	app.GET("/", controller.HomeContent)

	app.GET("/clients", controller.FindClients)

	app.Run()
}
