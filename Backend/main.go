package main

import (
	"backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// ⚙️ Habilita CORS para permitir frontend en otro puerto
	router.Use(cors.Default())

	// 🔐 Login
	router.POST("/login", func(c *gin.Context) {
		// Suponiendo que ya tenés esto implementado en controllers o directamente aquí
		// Podés reescribirlo como: controllers.Login(c)
		c.JSON(200, gin.H{"message": "Login dummy"})
	})

	// 📚 Actividades
	router.GET("/activities", controllers.GetActivities)
	router.GET("/activities/:id", controllers.GetActivityById)
	router.POST("/activities", controllers.AdminCreateActivity)
	router.PUT("/activities/:id", controllers.AdminUpdateActivity)
	router.DELETE("/activities/:id", controllers.AdminDeleteActivity)

	// 🚀 Start
	router.Run(":3000")
}
