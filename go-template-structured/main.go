package main

import (
	"go-template/routes"
	"go-template/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    // Inicializar conexión Mongo
    services.InitMongo()

    r := gin.Default()

    r.Use(cors.Default())

    // Registrar rutas
    routes.RegisterRoutes(r)

    r.Run(":8080")
}