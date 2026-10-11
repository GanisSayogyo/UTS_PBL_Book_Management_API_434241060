# Book Management API

REST API untuk mengelola data buku dan autentikasi pengguna. Project ini dikembangkan menggunakan Go, Fiber v2, PostgreSQL, JWT, bcrypt, dan Repository Pattern.

## Teknologi

- Go
- Fiber v2
- PostgreSQL
- pgx
- JWT
- bcrypt
- go-playground/validator

## Fitur

- Registrasi dan login pengguna.
- Autentikasi menggunakan JWT.
- Penyimpanan password menggunakan bcrypt.
- Role-based access control dengan role user dan admin.
- Melihat profil pengguna yang sedang login.
- CRUD data buku.
- Pencarian buku berdasarkan judul, penulis, atau ISBN.
- Pagination pada daftar buku.
- Validasi request dan penanganan error.

## Struktur Project

```text
app/
├── handler/
├── model/
├── repository/
└── service/
config/
database/
helper/
middleware/
migrations/
route/
main.go
```

## Persyaratan

Pastikan Go dan PostgreSQL sudah terpasang dan berjalan di komputer.

## Konfigurasi Environment

1. Salin file `.env.example` menjadi `.env`.
2. Sesuaikan konfigurasi database dan JWT di file `.env`.

Variabel environment yang digunakan:

- `APP_PORT`: port server API.
- `DB_HOST`: alamat host PostgreSQL.
- `DB_PORT`: port PostgreSQL.
- `DB_USER`: pengguna database.
- `DB_PASSWORD`: password database.
- `DB_NAME`: nama database.
- `JWT_SECRET`: secret untuk menandatangani JWT.

Jangan mengunggah file `.env` yang berisi kredensial ke repository.

## Database

Buat database PostgreSQL bernama `book_management`. Jalankan file migrasi berikut secara berurutan menggunakan PostgreSQL:

1. `migrations/001_create_users.sql`
2. `migrations/002_create_books.sql`

Migrasi membuat tabel `users` dan `books`.

## Menjalankan Project

Dari direktori utama project, jalankan:

```bash
go mod tidy
go run .
```

Server akan berjalan pada port yang ditentukan oleh `APP_PORT`, dengan port 3000 sebagai konfigurasi yang digunakan dalam pengujian.

## Endpoint API

Semua endpoint utama menggunakan prefix `/api/v1`.

| Method | Endpoint | Akses | Keterangan |
|---|---|---|---|
| POST | `/api/v1/auth/register` | Publik | Registrasi pengguna |
| POST | `/api/v1/auth/login` | Publik | Login pengguna |
| GET | `/api/v1/users/me` | User/Admin | Melihat profil sendiri |
| GET | `/api/v1/books` | User/Admin | Daftar buku |
| GET | `/api/v1/books/:id` | User/Admin | Detail buku |
| POST | `/api/v1/books` | Admin | Membuat buku |
| PUT | `/api/v1/books/:id` | Admin | Memperbarui seluruh data buku |
| PATCH | `/api/v1/books/:id` | Admin | Memperbarui sebagian data buku |
| DELETE | `/api/v1/books/:id` | Admin | Menghapus buku |
| GET | `/api/v1/health` | Publik | Memeriksa status API |

## Autentikasi dan Authorization

Untuk endpoint yang membutuhkan autentikasi, kirim token hasil login melalui header:

`Authorization: Bearer <token>`

Pengguna dengan role `user` dapat melihat data buku, sedangkan operasi pembuatan, pembaruan, dan penghapusan buku hanya dapat dilakukan oleh role `admin`.

Registrasi publik membuat akun dengan role `user`. Peningkatan role admin dilakukan melalui pengelolaan database oleh administrator.

## Query Parameter

Endpoint `GET /api/v1/books` mendukung:

- `page`: nomor halaman, dimulai dari 1.
- `limit`: jumlah data per halaman, maksimal 100.
- `search`: pencarian berdasarkan judul, penulis, atau ISBN.

Contoh:

`GET /api/v1/books?page=1&limit=10&search=Go`

Endpoint daftar buku memerlukan token autentikasi.

## Pengujian

Jalankan pengujian Go dengan perintah:

```bash
go test ./... -count=1
```

Pengujian manual endpoint dilakukan menggunakan PowerShell dan token JWT hasil login.

## Repository

GitHub: https://github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060
