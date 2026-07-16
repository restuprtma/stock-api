# Backend Dashboard Monitoring Stok Neraca Pangan Beras

Aplikasi ini merupakan MVP (Minimum Viable Product) Backend Dashboard untuk melakukan monitoring stok neraca pangan beras menggunakan Go. Menggunakan framework **Gin** untuk HTTP routing dan **GORM** dengan driver MySQL sebagai ORM.

> [!WARNING]
> **PERINGATAN PENTING:** Sebelum menjalankan aplikasi ini, Anda harus membuat database kosong di MySQL server lokal Anda bernama **`stok_beras_db`**.

---

## 🛠️ Persiapan Database

1. Buka MySQL client atau phpMyAdmin Anda.
2. Buat database baru dengan menjalankan perintah SQL berikut:
   ```sql
   CREATE DATABASE stok_beras_db;
   ```
3. Aplikasi diasumsikan terhubung ke MySQL di `127.0.0.1:3306` menggunakan user `root` tanpa password.

---

## 🚀 Cara Menjalankan Server secara Lokal

1. Pastikan Anda berada di direktori root project.
2. Unduh dan bersihkan dependensi dengan perintah:
   ```bash
   go mod tidy
   ```
3. Jalankan server menggunakan perintah:
   ```bash
   go run main.go
   ```
4. Server akan otomatis melakukan migrasi tabel (`stok_beras`) dan berjalan di **`http://localhost:8080`**.

---

## 📜 Kontrak API (OpenAPI v3.0.3)

Untuk integrasi Frontend, silakan rujuk ke spesifikasi di file [openapi.yaml](openapi.yaml).

### 1. Menambahkan Data Stok Beras
* **Endpoint:** `POST /api/v1/stok`
* **Content-Type:** `application/json`
* **Contoh Request Payload:**
  ```json
  {
    "tanggal": "2026-07-16",
    "jumlah_stok_ton": 1500.75,
    "lokasi_gudang": "Gudang Bulog Divre Indramayu",
    "catatan": "Stok beras premium jenis IR64 melimpah"
  }
  ```

### 2. Mendapatkan Semua Data Stok Beras
* **Endpoint:** `GET /api/v1/stok`
* **Contoh Response (200 OK):**
  ```json
  [
    {
      "id": 1,
      "tanggal": "2026-07-16T00:00:00Z",
      "jumlah_stok_ton": 1500.75,
      "lokasi_gudang": "Gudang Bulog Divre Indramayu",
      "catatan": "Stok beras premium jenis IR64 melimpah",
      "created_at": "2026-07-16T12:00:00Z",
      "updated_at": "2026-07-16T12:00:00Z"
    }
  ]
  ```

---

## 📊 Dataset Machine Learning Ready

Skema tabel `stok_beras` dirancang agar siap ditarik sebagai dataset Machine Learning:
* Tipe kolom `tanggal` didefinisikan sebagai tipe data `DATE` di MySQL demi integritas data temporal (memudahkan pre-processing deret waktu / forecasting).
* Nilai numeric `jumlah_stok_ton` menggunakan presisi ganda (`double`/`float64`).
* Ditambahkan indexing pada kolom `tanggal` dan `lokasi_gudang` untuk mempercepat agregasi data.
* Kolom audit `created_at` dan `updated_at` mempermudah analisis perubahan atau versi data dari waktu ke waktu.
