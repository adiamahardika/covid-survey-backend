# Covid Survey Backend

## Pendahuluan

Proyek ini adalah backend untuk aplikasi survei Covid-19 yang menggunakan Go, Gin, dan SQLite.

## Prasyarat

Pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/)
- [Git](https://git-scm.com/)

## Instalasi

1. Clone repositori ini:

   ```sh
   git clone <URL_REPOSITORY>
   cd covid-survey-backend
   ```

2. Unduh dependensi:
   ```sh
   go mod download
   ```

## Menjalankan Proyek

1. Jalankan server:

   ```sh
   go run main.go
   ```

2. Server akan berjalan di `http://localhost:5000`.

## Struktur Proyek

- `main.go`: File utama untuk menjalankan server.
- `routes/route.go`: File untuk mendefinisikan route API.
- `databases/`: Direktori untuk menyimpan database SQLite.

## API Endpoints

- **GET /surveys**: Mendapatkan semua survei.
- **POST /surveys**: Menambahkan survei baru.
- **GET /surveys/:id**: Mendapatkan survei berdasarkan ID.
- **PUT /surveys/:id**: Memperbarui survei berdasarkan ID.
- **DELETE /surveys/:id**: Menghapus survei berdasarkan ID.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
