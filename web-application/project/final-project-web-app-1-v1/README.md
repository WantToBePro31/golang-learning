# Final Project Web Application Backend

## Kanban App

### Description

Akhirnya sudah mencapai pembelajaran akhir di CAMP Ruangguru, ini adalah tugas akhir kalian yang wajib diselesaikan sebagai syarat kelulusan dari CAMP Ruangguru. Pada tugas ini, kalian akan membuat sebuah aplikasi web sederhana yang akan mengimplementasikan konsep REST API dan juga akan membangun sebuah aplikasi web yang akan mengimplementasikan konsep MVC.

Kalian diminta membuat Kanban App yang bisa digunakan untuk membuat sebuah _task_ (tugas) dan juga bisa mengelompokkan _task_ tersebut menjadi beberapa bagian.

Fitur yang harus dibuat baik dari sisi Rest API dan tampilan (_template_) web adalah sebagai berikut:

- Register user
- Login user
- Logout
- Create Category (Kategori)
- Delete Category
- Create Task (tugas)
- Update Task
- Delete Task
- Move Task (dari satu kategori ke kategori lain)

### Constraints

Terdapat beberapa _constraints_ yang harus kalian penuhi dalam mengerjakan tugas ini, yaitu:

- Aplikasi ini hanya dibuat menggunakan bahasa pemrograman Golang dan tidak menggunakan framework apapun.
- Terdapat 2 aplikasi yang berbeda di dalam satu _project_, yaitu aplikasi web dan aplikasi REST API.
- Aplikasi ini hanya diperkenankan menggunakan database PostgreSQL.
- Aplikasi ini hanya boleh di deploy ke fly.io dan akan diakses melalui internet.

### Todo

Kalian akan mengerjakan tugas ini dalam 3 tahap yang dibagi pada setiap minggu:

1. **Minggu 1**: Membuat aplikasi REST API, tugas dapat dilihat di file `phase1.md`
2. **Minggu 2**: Membuat aplikasi web (tampilan dan fungsionalitas), tugas dapat dilihat di file `phase2.md`
3. **Minggu 3**: Membuat dokumentasi dan deploy aplikasi ke fly.io, tugas dapat dilihat di file `phase3.md`

### Testing

Sebelum kalian menjalankan grader-cli test, pastikan kalian sudah mengubah environtment variable yaitu `"DATABASE_URL"` pada file `main.go` (line 45) dan `main_test.go` (line 26) sesuai dengan database kalian. Kalian cukup mengubah nilai dari "username", "password" dan "database_name" saja.

Contoh:

```go
os.Setenv("DATABASE_URL", "postgres://<username>:<password>@localhost:5432/<database_name>") // Ubah dengan credential database postgres di localhost.
```

## Demo application

kalian dapat mencobanya di [https://final-web-app.fly.dev](https://final-web-app.fly.dev)

<!-- TODO: masih perlu dibuat -->