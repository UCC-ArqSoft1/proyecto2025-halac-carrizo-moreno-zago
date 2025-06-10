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
	activity, err := services.GetActivityById(id)
	if err != nil || activity == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	c.JSON(http.StatusOK, activity)
}



// RegisterForActivity - Endpoint para inscribirse en una actividad
func RegisterForActivity(c *gin.Context) {
    id := c.Param("id")
    user := services.GetClientById("client1") // simulación, en producción usar JWT

    var req struct {
        DayOfWeek string `json:"day_of_week"`
    }
    if err := c.ShouldBindJSON(&req); err != nil || req.DayOfWeek == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Debe especificar el día de la semana"})
        return
    }

    activity, err := services.GetActivityById(id)
    if err != nil || activity == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
        return
    }

    // Validar que el día elegido exista en la actividad
    found := false
    for _, sch := range activity.Schedule {
        if sch.DayOfWeek == req.DayOfWeek {
            found = true
            break
        }
    }
    if !found {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Día de la semana inválido para esta actividad"})
        return
    }

    services.RegisterUserToActivity(user.ID, id, req.DayOfWeek)

    c.JSON(http.StatusOK, gin.H{
        "message":  "Registrado con éxito",
        "activity": activity,
        "user":     user,
        "day":      req.DayOfWeek,
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
