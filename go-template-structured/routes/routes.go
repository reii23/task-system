package routes

import (
    "github.com/gin-gonic/gin"
    "go-template/controllers"
)

func RegisterRoutes(router *gin.Engine) {
    router.GET("/ping", controllers.PingController)
}