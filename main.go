package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jeff-mendoza/integrator-dummy/business/controller"
)

func main() {
	router := gin.Default()
	router.GET("/ping", controller.Ping)
	router.GET("/webhook/:caller-id", controller.Check)
	router.GET("/webhook/event/:payment-intent-id", controller.FindEvent)
	router.POST("/webhook:caller-id", controller.Callback)

	router.GET("/config/:caller-id", controller.Find)
	router.POST("/config", controller.Create)

	router.Run("localhost:8080")
}
