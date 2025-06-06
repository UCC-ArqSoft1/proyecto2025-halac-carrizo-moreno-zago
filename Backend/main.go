package main

import (
	_ "backend/clients" // Para inicializar GORM + migraciones
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	services.SeedActivities()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

		cookie := &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, gin.H{"message": "login exitoso"})
	})

	// Rutas protegidas
	router.GET("/activities", middlewares.AuthMiddleware("admin", "socio"), controllers.GetActivities)
	router.GET("/activities/:id", controllers.GetActivityById) // <-- ¡AGREGADO!
	router.POST("/activities", middlewares.AuthMiddleware("admin"), controllers.AdminCreateActivity)
	router.GET("/user/activities", middlewares.AuthMiddleware("admin", "socio"), controllers.GetUserActivities)
	router.POST("/activities/:id/register", middlewares.AuthMiddleware("admin", "socio"), controllers.RegisterForActivity)


	// Ruta opcional /check-auth
	router.GET("/check-auth", middlewares.AuthMiddleware("admin", "socio"), func(c *gin.Context) {
		role := c.GetString("role")
		c.JSON(http.StatusOK, gin.H{"role": role})
	})

	router.Run(":3000")
}
