package controllers

import (
	"backend/domain"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUserActivities(c *gin.Context) {
    user := services.GetClientById("client1") // en producción: desde JWT

    activities := services.GetUserInscriptionsDetailed(user.ID)

    c.JSON(http.StatusOK, gin.H{
        "user":       user,
        "activities": activities,
    })
}

func GetActivities(c *gin.Context) {
    name := c.Query("name")
	var activities []domain.Activity

    if name != "" {
        // Asegurate de convertir a minúsculas si tu servicio espera eso
        activities = services.GetActivitiesByName(strings.ToLower(name))
    } else {
        activities = services.GetActivities()
    }

    c.JSON(http.StatusOK, activities)
}