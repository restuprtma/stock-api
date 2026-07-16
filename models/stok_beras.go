package models

import "time"

// StokBeras merupakan model database untuk mencatat persediaan stok beras.
// Dirancang secara ideal agar mudah ditarik sebagai dataset Machine Learning.
type StokBeras struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Tanggal       time.Time `gorm:"type:date;not null;index" json:"tanggal" binding:"required"` // Index ditambahkan untuk mempercepat query berdasarkan rentang waktu saat pre-processing dataset
	JumlahStokTon float64   `gorm:"type:double;not null" json:"jumlah_stok_ton" binding:"required"`
	LokasiGudang  string    `gorm:"type:varchar(255);not null;index" json:"lokasi_gudang" binding:"required"` // Index untuk pencarian/pengelompokkan berdasarkan gudang
	Catatan       string    `gorm:"type:text" json:"catatan"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CreateStokRequest merepresentasikan request payload JSON untuk membuat data stok.
type CreateStokRequest struct {
	Tanggal       string  `json:"tanggal" binding:"required"` // Format YYYY-MM-DD
	JumlahStokTon float64 `json:"jumlah_stok_ton" binding:"required"`
	LokasiGudang  string  `json:"lokasi_gudang" binding:"required"`
	Catatan       string  `json:"catatan"`
}
