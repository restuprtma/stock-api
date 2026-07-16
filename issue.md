---
name: Feature Request - MVP Backend Dashboard Monitoring Stok
about: Implementasi awal untuk Backend Dashboard Monitoring Stok Neraca Pangan Beras
title: "[FEATURE] MVP Backend Dashboard Monitoring Stok Neraca Pangan Beras"
labels: enhancement, backend
assignees: ''
---

## 📝 Deskripsi Tugas

Kita akan melakukan "vibe coding" untuk membangun MVP (Minimum Viable Product) Backend Dashboard Monitoring Stok Neraca Pangan Beras menggunakan bahasa pemrograman Go.

Environment Go (`go.mod`) dan dependensi library (seperti Gin, GORM, MySQL driver, CORS) **sudah disiapkan**. Tugas utama di issue ini adalah fokus pada penulisan kode sumber dan pembuatan dokumentasi.

## ✅ Task Checklist (Langkah-langkah Eksekusi)

Silakan centang checklist di bawah ini saat kamu menyelesaikan setiap bagian:

### 1. Buat Kontrak API (`openapi.yaml`)
- [ ] Buat file `openapi.yaml` (spesifikasi OpenAPI v3.0.3) di root folder proyek.
- [ ] Definisikan endpoint **POST `/api/v1/stok`**:
  - Menerima request body JSON (misalnya: tanggal, jumlah stok dalam ton, lokasi gudang, dan catatan).
- [ ] Definisikan endpoint **GET `/api/v1/stok`**:
  - Mengembalikan daftar semua data stok beras dalam format JSON.

### 2. Implementasi Kode Backend (`main.go`)
Tulis logika backend di dalam file `main.go` di root folder dengan ketentuan berikut:
- [ ] **Setup Database MySQL**: 
  - Gunakan `gorm.io/gorm`. 
  - Koneksikan menggunakan format DSN berikut: `root:@tcp(127.0.0.1:3306)/stok_beras_db?charset=utf8mb4&parseTime=True&loc=Local`. (Asumsikan username `root` tanpa password di localhost).
- [ ] **Model & AutoMigrate**: 
  - Buat struct Go untuk tabel stok beras (sesuaikan propertinya dengan kontrak OpenAPI).
  - Gunakan fungsi `AutoMigrate` dari GORM agar tabel terbuat secara otomatis.
- [ ] **Routing Gin**: 
  - Buat router menggunakan framework `github.com/gin-gonic/gin`.
- [ ] **Implementasi Handler**: 
  - Buat handler fungsional untuk endpoint `GET` dan `POST` `/api/v1/stok` yang memenuhi spesifikasi `openapi.yaml`.

### 3. Buat Dokumentasi (`README.md`)
Buat file `README.md` di root folder yang berisi:
- [ ] Instruksi lengkap tentang cara menjalankan server aplikasi (contohnya: `go run main.go`).
- [ ] **PERINGATAN PENTING**: Tambahkan peringatan tegas di bagian atas README yang menekankan bahwa user **wajib membuat database bernama `stok_beras_db` terlebih dahulu di MySQL** sebelum menjalankan aplikasi.

## ⚠️ Catatan Tambahan
Pastikan kode bisa di-compile dan dijalankan dengan baik tanpa error. Kerjakan dengan rapi, efisien, dan siap direview!
