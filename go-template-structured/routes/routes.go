package routes

import (
	"go-template/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.PingController)
	router.POST("/tasks", controllers.CrearTarea)
	router.GET("/tasks", controllers.ObtenerTareas)
	router.PUT("/tasks/:id/complete", controllers.CompletarTarea)
}
