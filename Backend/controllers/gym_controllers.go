package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGym - Endpoint para obtener la información del gimnasio
func GetGym(c *gin.Context) {
	gym := services.GetGym()
	c.JSON(http.StatusOK, gym)
}

// GetActivityById - Endpoint para obtener una actividad por su ID
func GetActivityById(c *gin.Context) {
	id := c.Param("id")
	activity := services.GetActivityById(id)
	c.JSON(http.StatusOK, activity)
}

// GetActivities - Endpoint para obtener todas las actividades disponibles
func GetActivities(c *gin.Context) {
	activities := services.GetActivities()
	c.JSON(http.StatusOK, activities)
}

// RegisterForActivity - Endpoint para inscribirse en una actividad
func RegisterForActivity(c *gin.Context) {
	id := c.Param("id")
	user := services.GetClientById("client1") // En el futuro: extraer desde el JWT
	activity := services.GetActivityById(id)

	// Simulación de inscripción
	c.JSON(http.StatusOK, gin.H{
		"message":  "Registrado con éxito",
		"activity": activity,
		"user":     user,
		"status":   "registered",
	})
}

// GetUserActivities - Endpoint para obtener las actividades de un usuario
func GetUserActivities(c *gin.Context) {
	user := services.GetClientById("client1")
	activities := []domain.Activity{
		services.GetActivityById("a1"), // En el futuro: recuperar actividades registradas del usuario
	}
	c.JSON(http.StatusOK, gin.H{
		"user":       user,
		"activities": activities,
	})
}

// AdminCreateActivity - Endpoint para que el administrador cree una actividad
func AdminCreateActivity(c *gin.Context) {
	var newActivity domain.Activity
	if err := c.ShouldBindJSON(&newActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.CreateActivity(newActivity)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Actividad creada con éxito",
		"activity": newActivity,
	})
}

// AdminUpdateActivity - Endpoint para que el administrador actualice una actividad
func AdminUpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var updatedActivity domain.Activity
	if err := c.ShouldBindJSON(&updatedActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.UpdateActivity(id, updatedActivity)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Actividad actualizada con éxito",
		"id":       id,
		"activity": updatedActivity,
	})
}

// AdminDeleteActivity - Endpoint para que el administrador elimine una actividad
func AdminDeleteActivity(c *gin.Context) {
	id := c.Param("id")
	services.DeleteActivity(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Actividad eliminada con éxito",
		"id":      id,
	})
}
