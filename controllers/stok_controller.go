package controllers

import (
	"net/http"
	"stock-api/config"
	"stock-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// GetStok menangani request GET /api/v1/stok untuk mendapatkan seluruh data stok beras.
func GetStok(c *gin.Context) {
	var stokList []models.StokBeras

	// Query semua data dari database diurutkan berdasarkan tanggal terbaru
	if err := config.DB.Order("tanggal desc").Find(&stokList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data dari database",
		})
		return
	}

	c.JSON(http.StatusOK, stokList)
}

// CreateStok menangani request POST /api/v1/stok untuk menyimpan data stok beras baru.
func CreateStok(c *gin.Context) {
	var req models.CreateStokRequest

	// Validasi JSON request payload sesuai schema OpenAPI
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parsing string tanggal (YYYY-MM-DD) ke tipe data time.Time
	parsedDate, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format tanggal tidak valid. Harus berupa YYYY-MM-DD.",
		})
		return
	}

	// Instansiasi model database
	stok := models.StokBeras{
		Tanggal:       parsedDate,
		JumlahStokTon: req.JumlahStokTon,
		LokasiGudang:  req.LokasiGudang,
		Catatan:       req.Catatan,
	}

	// Menyimpan ke database MySQL
	if err := config.DB.Create(&stok).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan data ke database",
		})
		return
	}

	c.JSON(http.StatusCreated, stok)
}
