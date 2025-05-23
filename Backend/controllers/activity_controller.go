package controllers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserActivities(c *gin.Context) {
	user := services.GetClientById("client1") // en producción: desde JWT

	activities := services.GetUserInscriptions(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"user":       user,
		"activities": activities,
	})
}
