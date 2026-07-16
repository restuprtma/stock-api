package config

import (
	"log"
	"stock-api/models"
	"time"
)

// SeedData memasukkan 6 baris data dummy ke tabel stok_beras jika tabel tersebut masih kosong.
func SeedData() {
	var count int64
	
	// Cek jumlah data yang ada di tabel stok_beras
	DB.Model(&models.StokBeras{}).Count(&count)

	if count > 0 {
		log.Println("Tabel stok_beras sudah memiliki data. Seeder dilewati.")
		return
	}

	log.Println("Tabel stok_beras kosong. Menjalankan database seeder...")

	// Membuat data dummy yang bervariasi untuk kebutuhan testing dan Machine Learning (time-series)
	dummyData := []models.StokBeras{
		{
			Tanggal:       time.Date(2026, 7, 10, 0, 0, 0, 0, time.Local),
			JumlahStokTon: 1200.50,
			LokasiGudang:  "Gudang Bulog Divre Indramayu",
			Catatan:       "Penerimaan beras premium hasil panen lokal",
		},
		{
			Tanggal:       time.Date(2026, 7, 11, 0, 0, 0, 0, time.Local),
			JumlahStokTon: 850.75,
			LokasiGudang:  "Gudang Bulog Divre Cianjur",
			Catatan:       "Cadangan pangan aman untuk sebulan ke depan",
		},
		{
			Tanggal:       time.Date(2026, 7, 12, 0, 0, 0, 0, time.Local),
			JumlahStokTon: 2100.00,
			LokasiGudang:  "Gudang Bulog Jakarta Utara",
			Catatan:       "Pemasukan beras impor kualitas premium asal Thailand",
		},
		{
			Tanggal:       time.Date(2026, 7, 13, 0, 0, 0, 0, time.Local),
			JumlahStokTon: 1450.25,
			LokasiGudang:  "Gudang Bulog Divre Karawang",
			Catatan:       "Distribusi beras bantuan pangan kemiskinan ekstrim",
		},
		{
			Tanggal:       time.Date(2026, 7, 14, 0, 0, 0, 0, time.Local),
			JumlahStokTon: 980.40,
			LokasiGudang:  "Gudang Bulog Jakarta Timur",
			Catatan:       "Persiapan pasokan untuk operasi pasar murah",
		},
		{
			Tanggal:       time.Date(2026, 7, 15, 0, 0, 0, 0, time.Local),
			JumlahStokTon: 1600.00,
			LokasiGudang:  "Gudang Bulog Divre Indramayu",
			Catatan:       "Tambahan stok beras medium jenis IR64",
		},
	}

	// Batch insert ke database
	if err := DB.Create(&dummyData).Error; err != nil {
		log.Printf("Gagal menjalankan database seeder: %v", err)
		return
	}

	log.Println("Database seeder berhasil dijalankan. 6 baris data dummy telah dibuat.")
}
