package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserActivities(c *gin.Context) {
	userID := c.GetInt("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	activities := services.GetUserInscriptionsDetailed(userID)

	c.JSON(http.StatusOK, gin.H{
		"user_id":    userID,
		"activities": activities,
	})
}

func GetActivities(c *gin.Context) {
	name := c.Query("name")
	var activities []domain.Activity

	if name != "" {
		// El servicio espera el nombre en minúsculas
		activities = services.GetActivitiesByName(strings.ToLower(name))
	} else {
		activities = services.GetActivities()
	}

	c.JSON(http.StatusOK, activities)
}