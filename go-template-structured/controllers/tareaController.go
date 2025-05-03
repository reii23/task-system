package controllers

import (
	"context"
	"go-template/models"
	"go-template/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// crear una nueva tara (con id, inicializada sin completar y la fecha en now)
func CrearTarea(c *gin.Context) {
	var tarea models.Tarea
	if error := c.ShouldBindJSON(&tarea); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error en formato de json"})
		return
	}

	tarea.ID = primitive.NewObjectID()
	tarea.Completed = false
	tarea.CreatedAt = time.Now()

	_, error := services.TaskCollection().InsertOne(context.Background(), tarea)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo crear la tarea "})

		return
	}

	c.JSON(http.StatusOK, tarea)

}

// obtener tareas trae todas las tareas
func ObtenerTareas(c *gin.Context) {
	cursor, error := services.TaskCollection().Find(context.Background(), bson.M{})
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error al traer las tareas"})

		return
	}
	var tareas []models.Tarea
	if error := cursor.All(context.Background(), &tareas); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudieron traer las tareas"})
		return
	}
	c.JSON(http.StatusOK, tareas)
}

// completar tarea por un id va a cambiar el estado completed de false a true
func CompletarTarea(c *gin.Context) {
	id := c.Param("id")
	objID, error := primitive.ObjectIDFromHex(id)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "formato del id invalido"})
		return
	}

	nuevo_estado := bson.M{"$set": bson.M{"completed": true}}

	_, error = services.TaskCollection().UpdateByID(context.Background(), objID, nuevo_estado)
	if error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo actualizar la tarea"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tarea": "tarea completada"})
}
