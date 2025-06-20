# Tiketin - Ticket Event Management API

Tiketin adalah aplikasi web berbasis rest api untuk pengelolaan event, tiket event, pemesanan tiket, dan fitur review.

## Fitur Utama

- **User Management**: Register, login, dan autentikasi JWT.
- **Event Management**: CRUD event, event type, dan ticket.
- **Order Management**: Pemesanan tiket, pembayaran, pembatalan, check-in, dan riwayat order.
- **Review**: User dapat memberikan review pada event yang diikuti.
- **Role-based Access**: Hak akses berdasarkan role (user, organizer, admin).
- **Migration**: Database migration dengan file SQL.

## Struktur Folder

```
.
├── controller/         # Handler untuk setiap endpoint
├── database/           # Koneksi dan migration database
├── helper/             # Helper & utilitas (JWT, validasi, dll)
├── middleware/         # Middleware (auth, role, dll)
├── model/              # Repository & service logic
├── public/             # Penyimpanan file yang bisa diakses public
├── routes/             # Routing API
├── structs/            # Struct data (DTO/model)
├── main.go             # Entry point aplikasi
└── ...
```

## Instalasi & Menjalankan

1. **Clone repository**
   ```sh
   git clone <repo-url>
   cd Tiketin
   ```

2. **Atur konfigurasi database** (optional jika butuh migrasi baru maka perlu setup)
   - Edit file `dbconfig.yml` atau gunakan environment variable sesuai kebutuhan.

3. **Jalankan migration**
   ```sh
   go run main.go
   ```
   (Migration akan berjalan otomatis saat aplikasi start.)

4. **Jalankan aplikasi**
   ```sh
   go run main.go
   ```
   Akses API di `http://localhost:8080`

## Contoh Endpoint

- **Register User:** `POST /api/users/register`
- **Login User:** `POST /api/users/login`
- **Get All Event:** `GET /api/events/list`
- **Create Event:** `POST /api/events/create`
- **Order Ticket:** `POST /api/orders/create`
- **Checkin Ticket:** `GET /api/orders/checkin/ticket?orderId=7&orderItemId=8&ticketId=1`
- **Review Event:** `POST /api/events/:event_id/reviews/create`

## Environment Variable
config ada di config/config.json
- `Database.Host`
- `Database.Port`
- `Database.User`
- `Database.Pass`
- `Database.DbName`
- `App.Port`
- `App.BaseUrl` (url yang digunakan untuk qr code)
- `JWT.SignatureKey`
- `JWT.Issuer`

## Tools & Library

- [Gin](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org/)
- `bcrypt` untuk hash password
- [viper](https://github.com/spf13/viper)
- [go-playground/validator](https://github.com/go-playground/validator)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [sql-migrate](https://github.com/rubenv/sql-migrate)
- [qr-code](https://github.com/skip2/go-qrcode)
- [swag](https://github.com/swaggo/swag)
- dan lainnya

## Docs Api (Swagger)
https://{host}/swagger/index.html

**Lisensi:** MIT
