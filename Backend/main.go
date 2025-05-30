package main

import (
	"backend/controllers"
	"backend/services"
	

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	router := gin.Default()

	// ⚙️ Habilita CORS para permitir frontend en otro puerto
	router.Use(cors.Default())

	// 🔐 Login
	router.POST("/login", func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}
	
		token, err := services.Login(credentials.Username, credentials.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
	

	// 📚 Actividades
	router.GET("/activities", controllers.GetActivities)
	router.GET("/activities/:id", controllers.GetActivityById)
	router.POST("/activities", controllers.AdminCreateActivity)
	router.PUT("/activities/:id", controllers.AdminUpdateActivity)
	router.DELETE("/activities/:id", controllers.AdminDeleteActivity)
	router.POST("/activities/:id/register", controllers.RegisterForActivity)
	router.GET("/user/activities", controllers.GetUserActivities)



	// 🚀 Start
	router.Run(":3000")
}