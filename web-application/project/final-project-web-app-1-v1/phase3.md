# Week 3

## Deployment

### Description

Jika proses pengetesan aplikasi sudah selesai, maka saatnya kita melakukan proses deployment. Deployment adalah proses untuk mengupload aplikasi yang sudah kita buat ke server. Proses deployment ini akan dilakukan dengan menggunakan layanan [fly.io](https://fly.io/).

### Instruction

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

  ? Choose an app name (leave blank to generate one):       # <nama-campid>
  ? Select Organization: ruangguru                          # pilih organisasi ruangguru
  ? Choose a region for deployment: singapore               # Pilih server region singapore` 
  Created app rg-app in organization personal
  Wrote config file fly.toml
  ? Would you like to set up a Postgresql database now?     # Yes -> Konfigurasi database postgres fly.io
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
  ? Would you like to set up an Upstash Redis database now? No  # harus pilih No
  ? Would you like to deploy now?                               # Yes -> Langsung melakukan deploy

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

    Nama aplikasi harus sesuai format yaitu `<nama>-final-project-<campid>` contoh: **aditira-final-project-BE00000** jika mengikuti langkah diatas, maka `postgres-app-name` akan menjadi **aditira-final-project-BE00000-db**

    > **Note**: apabila muncul pesan error `Error Validation failed: Name has already been taken`, ubah nama app menjadi <nama>-final-project-<campid>-<number> contoh: **aditira-final-project-BE00000-1** atau **aditira-final-project-BE00000-1** (tambahkan angka setelah camp id sehingga tidak muncul pesan error tersebut lagi)

- Lakukan perintah `flyctl secrets set DATABASE_URL=postgres://postgres:<password>@<nama-campid>-db.internal:5432` untuk mengatur konfigurasi database yang sudah dibuat pada langkah sebelumnya:

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
    return "https://final-web-app.fly.dev"
  }
  ```

- Submit menggunakan grader-cli : `grader-cli submit`
  > **Penting**: Jangan melampirkan url <https://aditira-2022.fly.dev> pada saat sumbit, jika iya maka penilaian deploy **0**. Ganti dengan url deploy masing-masing sesuai dengan format nama aplikasi yang ditentukan!

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah _environtment variable_ yaitu **"DATABASE_URL"** pada file **`main.go`** (line 45) dan **`main_test.go`** (line 27) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

Contoh:

```go
os.Setenv("DATABASE_URL", "postgres://<username>:<password>@localhost:5432/<database_name>") // Ubah dengan credential database postgres di localhost.
```

Jika hendak melakukan deploy ke `fly.io` **hapus code di atas**

> **Note:** Aplikasi yang di deploy akan dihapus secara berkala setiap hari (1x24 jam). Maka dari itu setelah Anda berhasil deploy aplikasi harap  segera jalankan perintah grader-cli submit agar pekerjaan Anda bisa dinilai.
