---
name: Setup Unit Testing Backend
about: Membuat struktur unit testing menggunakan SQLite in-memory dan test endpoint
title: "[TESTING] Setup Unit Testing untuk Controller Backend"
labels: enhancement, testing
assignees: ''
---

## 📝 Deskripsi Tugas

Kita perlu membangun struktur unit testing yang kokoh untuk proyek **MVP Backend Dashboard Monitoring Stok Neraca Pangan Beras** (`stock-api`).

Project ini sudah menerapkan arsitektur modular (`config/`, `models/`, `controllers/`, `routes/`). Tugasmu adalah menambahkan library testing dan membuat unit test untuk memastikan endpoint berjalan dengan baik.

## ✅ Task Checklist (Langkah-langkah Eksekusi)

### 1. Setup Dependency
- [ ] Instal package *test assertion*: `go get github.com/stretchr/testify/assert`
- [ ] Instal driver database *in-memory*: `go get gorm.io/driver/sqlite`
*(Catatan: Driver SQLite ini hanya digunakan khusus untuk environment testing agar tidak mengotori database utama MySQL).*

### 2. Buat File Konfigurasi Test
- [ ] Buat file helper untuk testing (misalnya `tests/setup_test.go` atau di dalam folder `config/`).
- [ ] Inisialisasi koneksi database testing menggunakan SQLite in-memory: `gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})`.
- [ ] Lakukan `AutoMigrate` pada database in-memory tersebut menggunakan model `models.StokBeras`.
- [ ] Buat fungsi untuk mengembalikan `*gin.Engine` dalam mode `TestMode` (`gin.SetMode(gin.TestMode)`).
- [ ] Pastikan global `config.DB` tergantikan dengan koneksi database in-memory ini saat testing berjalan.

### 3. Buat Unit Test untuk Endpoint GET dan POST
- [ ] Buat file unit test, misalnya `controllers/stok_controller_test.go`.
- [ ] Gunakan package bawaan `net/http/httptest` untuk melakukan mock HTTP request.
- [ ] Berikan komentar yang rapi pada setiap test function dengan pola: `// Setup` -> `// Action` -> `// Assert`.

**Test Case POST (`POST /api/v1/stok`):**
- [ ] **Skenario Sukses:** Kirim payload JSON yang valid. Lakukan *assert* bahwa API mengembalikan `HTTP 201 Created` dan responsnya memiliki data yang benar.
- [ ] **Skenario Gagal:** Kirim payload yang tidak valid (misal: format tanggal disengaja salah atau field required ada yang kosong). Lakukan *assert* bahwa API mengembalikan `HTTP 400 Bad Request`.

**Test Case GET (`GET /api/v1/stok`):**
- [ ] **Skenario Sukses:** Insert data dummy via code ke in-memory DB di tahap Setup (jika perlu). Kirim request GET, lalu lakukan *assert* bahwa API mengembalikan `HTTP 200 OK`.
- [ ] Verifikasi bahwa response body berupa array JSON dan berisikan data yang sesuai.

### 4. Instruksi Menjalankan Test
- [ ] Tuliskan semua kodenya dengan benar dan bersih.
- [ ] Pastikan seluruh test lulus saat dijalankan dengan perintah `go test -v ./...`.

## ⚠️ Catatan Tambahan
Jalankan test lokal sebelum membuat Pull Request untuk memastikan tidak ada *test* yang bocor (*leak state*) ke *test* lain.
