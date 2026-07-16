package main

import (
	"log"
	"stock-api/config"
	"stock-api/routes"
)

func main() {
	// Inisialisasi koneksi database, AutoMigrate, dan database Seeder
	config.InitDB()

	// Inisialisasi Gin Engine & Routes
	r := routes.SetupRouter()

	// Menjalankan server di port 8080
	log.Println("Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
