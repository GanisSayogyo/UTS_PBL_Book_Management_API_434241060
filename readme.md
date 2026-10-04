# Book Management API

Backend REST API untuk mengelola data buku dan autentikasi pengguna.

Project ini dibuat menggunakan Go, Fiber v2, PostgreSQL, JWT, bcrypt, dan Repository Pattern.

## Teknologi

- Go
- Fiber v2
- PostgreSQL
- pgx
- JWT
- bcrypt
- go-playground/validator

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