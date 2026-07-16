package tests

import (
	"stock-api/config"
	"stock-api/models"
	"stock-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB menginisialisasi SQLite in-memory dan AutoMigrate untuk environment testing.
func SetupTestDB() {
	var err error
	// Override global config.DB dengan SQLite in-memory
	config.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Gagal setup test database: " + err.Error())
	}

	// AutoMigrate khusus test
	err = config.DB.AutoMigrate(&models.StokBeras{})
	if err != nil {
		panic("Gagal automigrate test database: " + err.Error())
	}

	// Bersihkan state antar test
	config.DB.Exec("DELETE FROM stok_beras")
}

// SetupTestRouter mengatur mode test Gin dan mengembalikan engine router yang telah dikonfigurasi.
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	// Kita bisa langsung memanggil SetupRouter dari package routes
	return routes.SetupRouter()
}
