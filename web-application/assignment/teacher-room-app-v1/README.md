# Web Application

## Tacher Room App with GORM

### Description

Terdapat aplikasi sederhana bernama **Teacher Room** yang mengimplementasikan **CRUD** biasa menggunakan **GORM** pada tabel **teachers** yang memiliki kolom `id`, `created_at`, `updated_at`, `deleted_at`, `name`, `field_of_study` dan `age`. Disini kamu diminta untuk melengkapi layer **repository** pada file `app/repository/teacher.go` yang memiliki method:

- `AddTeacher(teacher model.Teacher)`: menerima parameter bertipe `model.Teacher` dan berfungsi menyimpan data sesuai parameter tersebut ke tabel `teachers`.
- `ReadTeacher()`: menampilkan semua data pada tabel `teachers` di database dalam bentuk model `[]model.ViewTeacher`, tampilkan kolom `name`, `field_of_study` dan `age` pada tabel tersebut dengan kondisi kolom `deleted_at` adalah `NULL`.
- `UpdateName(id uint, name string)`: menerima parameter **id** dengan _type_ `uint` dan **name** dengan _type_ `string` dan berfungsi untuk mengubah data `teachers` column `name` dari parameter sesuai dengan dengan `id` yang diterima. mengubah kolom `name` pada tabel sesuai dengan perameter yang diterima.
- `DeleteTeacher(id uint)`: menerima parameter **id** dengan _type_ `uint` dan berfungsi untuk menghapus data `teachers` sesuai dengan dengan `id` yang diterima.

Setelah kita melengkapi method handler di atas, maka kita wajib melakukan deploy ke platform server **[fly.io](https://fly.io/)** dengan ketentuan sebagai berikut.

### Constraints

Berikut adalah ketentuan yang harus dilakukan untuk melakukan deploy aplikasi ini ke **[fly.io](https://fly.io/)**:

- Hapus atau comment baris `os.Setenv` pada file `main.go`:
  
  ```go
  os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/kampusmerdeka")
  ```

- Lakukan perintah `flyctl launch` di root directory untuk membuat konfigurasi aplikasi pada project ini (konfigurasi applikasi dan database):

  ```bash
  > flyctl launch
  Creating app in <your local path>
  Scanning source code
  Detected a Go app
  Using the following build configuration:
          Builder: paketobuildpacks/builder:base
          Buildpacks: gcr.io/paketo-buildpacks/go

  ? Choose an app name (leave blank to generate one): # <nama-campid>
  ? Select Organization: ruangguru # pilih organisasi ruangguru
  ? Choose a region for deployment: singapore # Pilih server region singapore` 
  Created app rg-app in organization personal
  Wrote config file fly.toml
  ? Would you like to set up a Postgresql database now? # Yes -> Konfigurasi database postgres fly.io
  ? Select configuration: Development - Single node, 1x shared CPU, 256MB RAM, 1GB disk
  Creating postgres cluster in organization personal
  Creating app...
  Setting secrets on app <nama-campid>-db...
  ...
  Postgres cluster <nama-campid>-db created
    Username:    postgres
    Password:    eTFLtqsqbIgF7xo
    Hostname:    <nama-campid>.internal
    Proxy port:  5432
    Postgres port:  5433
    Connection string: postgres://postgres:eTFLtqsqbIgF7xo@<nama-campid>-db.internal:5432
  ...
  ? Would you like to set up an Upstash Redis database now? No # harus pilih No
  ? Would you like to deploy now? # Yes -> Langsung melakukan deploy

  ==> Verifying app config
  --> Verified app config
  ==> Building image
  Remote builder fly-builder-black-wind-90 ready
  ==> Building image with Buildpacks
  ...
  # tunggu proses deployment sampai selesai.
  ...
   1 desired, 1 placed, 1 healthy, 0 unhealthy [health checks: 1 total, 1 passing]
  --> v0 deployed successfully
  ```

  > **Note**: jangan lupa, nama aplikasi harus sesuai format yaitu `<nama-campid>` contoh: **aditira-2022** jika mengikuti langkah diatas, maka `postgres-app-name` akan menjadi **aditira-2022-db**

- Setelah melakukan deploy, kita harus set env variable `DATABASE_URL` di server Fly dengan perintah:

  ```bash
  > flyctl secrets set DATABASE_URL=postgres://<username>:<password>@<postgres-app-name>.internal:5432/<database_name>
  ```

- Check apakah env variable `DATABASE_URL` sudah berhasil di pasang dengan perintah:

  ```bash
  > flyctl secrets list
  ```

- Jika proses deploy berhasil dan success, maka kita bisa buka dengan perintah `flyctl open`. Ini akan membuka aplikasi yang telah berhasil kita deploy melaui browser default.
- Jika aplikasi sudah terbuka, maka copy url dari aplikasinya dan tambahkan pada fungsi `FlyURL()` di file `main.go`.

  ```go
  func FlyURL() string {
    return "" // ubah dengan url dari aplikasi yang telah di deploy!
  }
  ```

  Contoh

  ```go
  func FlyURL() string {
    return "https://aditira-2022.fly.dev"
  }
  ```

  > **Penting**: Jangan melampirkan url <https://aditira-2022.fly.dev> pada saat sumbit, jika iya maka penilaian deploy **0**. Ganti dengan url deploy masing-masing sesuai dengan format nama aplikasi yang ditentukan!

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah _environtment variable_ yaitu **"DATABASE_URL"** pada file **`main.go`** (line 19) dan **`main_test.go`** (line 21) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

Contoh:

```go
os.Setenv("DATABASE_URL", "postgres://<username>:<password>@localhost:5432/<database_name>") // Ubah dengan credential database postgres di localhost.
```

Jika hendak melakukan deploy ke `fly.io` **hapus code di atas** dan lakukan setup env dengan `flyctl`:

```bash
flyctl secrets set DATABASE_URL=postgres://<username>:<password>@<postgres-app-name>.internal:5432/<database_name>
```

> **Note:** Applikasi yang di deploy akan dihapus secara otomatis setiap hari (1/24 jam).
