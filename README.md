# TaskFlow

## Tentang aplikasi ini

TaskFlow adalah aplikasi manajemen tugas (to-do list) berbasis web. Setiap orang yang pakai aplikasi ini wajib daftar akun dan login dulu sebelum bisa membuat atau mengelola daftar tugas hariannya. Data tiap akun terpisah sepenuhnya, jadi satu user tidak bisa melihat atau mengubah tugas milik user lain.

Fitur utamanya:

- Registrasi dan login akun, memakai autentikasi berbasis token (JWT), password disimpan dalam bentuk hash, bukan teks biasa.
- Tambah tugas baru, lengkap dengan judul, deskripsi, tingkat prioritas (Low, Medium, High), dan tenggat waktu.
- Lihat daftar tugas, bisa difilter berdasarkan status selesai atau belum, dan diurutkan berdasarkan tenggat waktu.
- Tandai tugas sebagai selesai atau ubah detailnya kapan saja.
- Hapus tugas yang sudah tidak diperlukan.

Aplikasi ini terdiri dari dua bagian yang saling terhubung, backend (server API di folder `taskflow-backend`) dan frontend (tampilan web di folder `taskflow-frontend`). Keduanya perlu dijalankan bersamaan supaya aplikasi berfungsi penuh.

## Cara menjalankan aplikasi

### 1. Jalankan backend

Buka terminal, masuk ke folder backend, lalu jalankan:

```
cd taskflow-backend
go mod download
go run ./cmd/api
```

Tunggu sampai muncul tulisan `TaskFlow API listening on :8080`. Biarkan terminal ini tetap terbuka selama Anda memakai aplikasi.

Konfigurasi backend bisa diatur lewat file `.env` (contoh ada di `.env.example`):

- `PORT`, port server, default 8080
- `DATABASE_DSN`, lokasi file database SQLite, dibuat otomatis kalau belum ada
- `JWT_SECRET`, kunci rahasia untuk token login, wajib diganti kalau aplikasi mau dipakai serius di luar komputer sendiri

### 2. Jalankan frontend

Buka terminal baru (jangan tutup terminal backend), masuk ke folder frontend, lalu jalankan:

```
cd taskflow-frontend
npm install
cp .env.local.example .env.local
npm run dev
```

Tunggu sampai muncul tulisan `Local: http://localhost:3000`.

## Cara menggunakan aplikasi

1. Buka `http://localhost:3000` di browser.
2. Klik "Create account" untuk daftar. Isi nama, email, dan password minimal 8 karakter.
3. Setelah daftar berhasil, Anda otomatis masuk ke halaman dashboard.
4. Klik tombol "+ New task" untuk menambah tugas baru. Isi judul, deskripsi, pilih prioritas, dan atur tenggat waktu kalau perlu.
5. Klik lingkaran di sebelah kiri tugas untuk menandainya selesai.
6. Gunakan tab All, Pending, atau Completed untuk menyaring daftar tugas yang ditampilkan.
7. Arahkan kursor ke sebuah tugas untuk memunculkan tombol hapus, lalu klik untuk menghapusnya.
8. Kalau lain waktu ingin login lagi, kembali ke `http://localhost:3000`, klik "Sign in", masukkan email dan password yang sama.

Catatan, kedua server (backend port 8080 dan frontend port 3000) harus tetap jalan bersamaan selama Anda memakai aplikasi. Kalau salah satu terminal ditutup, aplikasi berhenti berfungsi.

## Endpoint API

| Method | Path | Perlu login |
|---|---|---|
| POST | /api/v1/auth/register | tidak |
| POST | /api/v1/auth/login | tidak |
| GET | /api/v1/todos | ya |
| POST | /api/v1/todos | ya |
| PUT | /api/v1/todos/:id | ya |
| DELETE | /api/v1/todos/:id | ya |

Setiap endpoint todo membaca header `Authorization: Bearer <token>` dan hanya mengembalikan data milik user yang login. Semua query database difilter dengan `user_id` untuk mencegah user mengakses data user lain.

## Catatan implementasi

- Database memakai SQLite, file lokal untuk kemudahan setup. Skema dan query sudah kompatibel pola dengan PostgreSQL/MySQL, kalau nanti mau migrasi tinggal ganti driver dan DSN di `internal/repository/db.go`.
- Password di-hash dengan bcrypt, cost factor 10.
- Token JWT berlaku 24 jam sejak login.
- Proteksi rute ada di dua sisi, backend (`internal/delivery/http/middleware/auth_middleware.go`) dan frontend (`src/app/(dashboard)/layout.tsx`, redirect ke halaman login kalau tidak ada token tersimpan).
