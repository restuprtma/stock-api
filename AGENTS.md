# AGENTS.md

## konteks
Proyek ini adalah **MVP Backend Dashboard Monitoring Stok Neraca Pangan Beras** yang dibangun menggunakan bahasa pemrograman **Go (Golang)**. 
- Framework routing yang digunakan adalah **Gin Web Framework**.
- Database yang digunakan adalah **MySQL** dengan **GORM** sebagai ORM-nya.
- Seluruh dependensi library eksternal (Gin, GORM, MySQL Driver, CORS Middleware) sudah terpasang dan dikelola di dalam file `go.mod`.

## kontrak diatas segalanya
- Kontrak API yang didefinisikan di dalam file `openapi.yaml` di root folder proyek merupakan **Single Source of Truth (Satu-satunya Sumber Kebenaran)**.
- AI Agent dilarang keras memodifikasi, menambah, atau mengurangi endpoint, nama kolom JSON, tipe data request, maupun format response secara sepihak.
- Perubahan kontrak harus diperbarui dan disepakati di dalam `openapi.yaml` terlebih dahulu sebelum melakukan penulisan kode di `main.go`.

## verifikasi sebelum menyatakan selesai
Sebelum menyatakan tugas selesai, AI Agent harus memastikan seluruh langkah verifikasi berikut lolos:
1. **Verifikasi Sintaksis (Build):** Jalankan perintah `go build` di root folder untuk memastikan kode Go terkompilasi dengan sukses tanpa error sintaksis.
2. **Koneksi Database:** Pastikan koneksi ke MySQL berhasil dan proses `AutoMigrate` berjalan lancar tanpa mengalami panic/crash di konsol.
3. **Validasi Endpoint (API Test):** Lakukan pengujian endpoint menggunakan curl atau HTTP client untuk memastikan:
   - `POST /api/v1/stok` dapat menyimpan data dengan format request yang sesuai dan mengembalikan status code `201 Created`.
   - `GET /api/v1/stok` mengembalikan seluruh daftar data dalam format array JSON dengan status code `200 OK`.

## batasan
- **Port Layanan:** Backend harus berjalan di port `:8080` (`http://localhost:8080`) secara default.
- **Konfigurasi CORS:** Middleware CORS hanya boleh mengizinkan akses dari origin `http://localhost:3000`. Jangan gunakan wildcard `*` untuk origin.
- **Kesiapan Dataset ML:** Skema tabel `stok_beras` harus dipertahankan agar memiliki struktur ideal untuk diolah menjadi dataset Machine Learning (analisis tren/prediksi):
  - Kolom `tanggal` wajib menggunakan tipe data `DATE` di MySQL untuk memudahkan analisis deret waktu.
  - Jumlah stok wajib memiliki tipe numerik desimal presisi tinggi (`double` atau `float64`).
  - Tambahkan indeks pada kolom penting (`tanggal` dan `lokasi_gudang`) untuk kecepatan kueri saat pre-processing data.
  - Wajib menyertakan metadata temporal (`created_at` dan `updated_at`).
- **Asumsi Lingkungan:** Koneksi database lokal diasumsikan menggunakan DSN `root:@tcp(127.0.0.1:3306)/stok_beras_db?charset=utf8mb4&parseTime=True&loc=Local`. AI Agent wajib mengingatkan pengguna untuk membuat database `stok_beras_db` terlebih dahulu sebelum menjalankan server.
