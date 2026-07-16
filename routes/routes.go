package routes

import (
	"stock-api/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi Gin Engine, mengkonfigurasi CORS, dan mendaftarkan route API.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Konfigurasi CORS agar menerima request dari http://localhost:3000
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Grouping routing API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/stok", controllers.GetStok)
		v1.POST("/stok", controllers.CreateStok)
	}

	return r
}
