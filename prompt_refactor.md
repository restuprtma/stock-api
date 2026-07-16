---
name: Refactor Backend & Seeder Dummy Data
about: Refactoring struktur folder main.go dan penambahan dummy data
title: "[REFACTOR] Struktur Folder Modular & Seeder Dummy Data"
labels: enhancement, refactor
assignees: ''
---

## 📝 Deskripsi Tugas

Kita perlu melanjutkan pengembangan proyek **MVP Backend Dashboard Monitoring Stok Neraca Pangan Beras**. 

Saat ini, seluruh kode backend masih berada di dalam satu file `main.go`. Tugas utamamu adalah merapikan (refactor) kode tersebut ke dalam struktur folder yang bersih dan standar, serta menambahkan fitur data dummy (seeder) agar database langsung terisi untuk keperluan testing.

**PENTING**: Sebelum mulai, baca terlebih dahulu file `AGENTS.md` untuk memahami batasan, konteks, dan aturan utama (terutama mengenai kepatuhan pada `openapi.yaml`).

## ✅ Task Checklist (Langkah-langkah Eksekusi)

### 1. Refactor Struktur Folder
Pecah kode yang ada di `main.go` ke dalam struktur folder modular berikut:
- [ ] **`config/`** -> Buat file `database.go` untuk menangani koneksi MySQL (DSN) dan inisialisasi GORM.
- [ ] **`models/`** -> Buat file `stok_beras.go` untuk memindahkan definisi struct `StokBeras` dan request payload.
- [ ] **`controllers/` (atau `handlers/`)** -> Buat file `stok_controller.go` untuk memindahkan logika handler fungsional (`getStok` dan `createStok`).
- [ ] **`routes/`** -> Buat file `routes.go` untuk mengatur inisialisasi Gin router, middleware CORS, dan pendaftaran endpoint `/api/v1/stok`.
- [ ] **`main.go`** -> Bersihkan `main.go` sehingga hanya memanggil inisialisasi config, routes, dan menjalankan server.

### 2. Buat Seeder Data Dummy
- [ ] Buat logika seeder: "Jika tabel `stok_beras` kosong setelah `AutoMigrate`, maka insert data dummy". Logika ini bisa ditempatkan di `config/database.go` atau file khusus `config/seeder.go`.
- [ ] Masukkan minimal 5-10 baris data dummy stok beras yang realistis (dengan variasi `tanggal`, `jumlah_stok_ton`, `lokasi_gudang`, `catatan`).

### 3. Verifikasi & Aturan Keselamatan
- [ ] **Dilarang** mengubah kontrak `openapi.yaml` atau struktur tabel (untuk Machine Learning).
- [ ] Pastikan package import antar folder valid (menggunakan module `stock-api`).
- [ ] Pastikan hasil refactor bisa dikompilasi tanpa error (`go build`).
- [ ] Lakukan pengetesan manual (via Postman/curl) untuk memastikan endpoint GET dan POST tetap berfungsi sama persis seperti sebelumnya.

## ⚠️ Catatan Tambahan
Pastikan seluruh perubahan ini berjalan lancar di lokal (terhubung ke `stok_beras_db`) sebelum membuat Pull Request!
