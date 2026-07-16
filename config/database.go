package config

import (
	"log"
	"stock-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah instance global database
var DB *gorm.DB

// InitDB menginisialisasi koneksi MySQL GORM, menjalankan AutoMigrate, dan memanggil seeder.
func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/stok_beras_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// Membuka koneksi database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	log.Println("Koneksi database berhasil dilakukan.")

	// Menjalankan AutoMigrate untuk model StokBeras
	err = DB.AutoMigrate(&models.StokBeras{})
	if err != nil {
		log.Fatalf("Gagal melakukan AutoMigrate: %v", err)
	}

	log.Println("AutoMigrate tabel stok_beras berhasil.")

	// Jalankan seeder data dummy
	SeedData()
}
