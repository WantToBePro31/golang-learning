# Week 1

## Rest API

### Database Model and Schema

![db-relation-model](./assets/markdown/db-relation-model.png)

Aplikasi ini memiliki 3 tabel utama, yaitu `users`, `categories`, dan `tasks`. Tabel `users` digunakan untuk menyimpan data-data user, tabel `categories` digunakan untuk menyimpan data-data kategori, dan tabel `tasks` digunakan untuk menyimpan data-data task.

Tabel `users` dapat memiliki banyak kategori, dan tabel `categories` dapat memiliki banyak task. Tabel `users` dan `categories` memiliki relasi one-to-many, sedangkan tabel `categories` dan `tasks` juga memiliki relasi one-to-many. Di dalam table `tasks` terdapat kolom `user_id` yang digunakan untuk menyimpan id dari user yang membuat task tersebut, dan kolom `category_id` yang digunakan untuk menyimpan id dari kategori yang dimiliki task tersebut.

### Instruction

Berikut adalah hal-hal yang harus diperhatikan dalam mengerjakan Rest API Kanban App:

📁 **repository**

Ini adalah fungsi yang berinteraksi dengan database Postgres menggunakan GORM:

> Warning : abaikan code yang tidak berhubungan dengan instruksi di bawah ini pada folder `repository`

- **users**
  - method `GetUserByID`:  menerima paramter `context.Context` dan `id` bertipe `int`. Fungsi ini akan mengembalikan data user berdasarkan `id` yang diberikan.
    - Jika terjadi error saat mengambil data user, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika user tidak ditemukan, maka fungsi ini akan mengembalikan `entity.User{}` dan `nil`.
    - Jika user ditemukan, maka fungsi ini akan mengembalikan `entity.User` dan `nil`.
  - method `GetUserByEmail`: menerima paramter `context.Context` dan `email` bertipe `string`. Fungsi ini akan mengembalikan data user berdasarkan `email` yang diberikan.
    - Jika terjadi error saat mengambil data user, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika user tidak ditemukan, maka fungsi ini akan mengembalikan `entity.User{}` dan `nil`.
    - Jika user ditemukan, maka fungsi ini akan mengembalikan `entity.User` dan `nil`.
  - method `CreateUser`: menerima paramter `context.Context` dan `user` bertipe `User`. Fungsi ini akan membuat data user baru berdasarkan `user` yang diberikan.
    - Jika terjadi error saat membuat data user, maka fungsi ini akan mengembalikan struct kosong dan `error`.
    - Jika user berhasil dibuat, maka fungsi ini akan mengembalikan hasil create `entity.User` dan `nil`.
  - method `UpdateUser`: menerima paramter `context.Context` dan `user` bertipe `User`. Fungsi ini akan mengupdate data user berdasarkan `user` yang diberikan.
    - Jika terjadi error saat mengupdate data user, maka fungsi ini akan mengembalikan struct kosong dan `error`.
    - Jika user berhasil diupdate, maka fungsi ini akan mengembalikan struct yang sudah diperbarui dan `nil`.
  - method `DeleteUser`: menerima paramter `context.Context` dan `id` bertipe `int`. Fungsi ini akan menghapus data user berdasarkan `id` yang diberikan.
    - Jika terjadi error saat menghapus data user, maka fungsi ini akan mengembalikan `error`.
    - Jika user berhasil dihapus, maka fungsi ini akan mengembalikan `nil`.

- **category**
  - method `GetCategoriesByUserId` : menerima paramter `context.Context` dan `id` bertipe `int` (_user id_). Fungsi ini akan mengembalikan data kategori berdasarkan user id di parameter `id` yang diberikan.
    - Jika terjadi error saat mengambil data kategori, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika kategori tidak ditemukan, maka fungsi ini akan mengembalikan `[]entity.Category{}` dan `nil`.
    - Jika kategori ditemukan, maka fungsi ini akan mengembalikan `[]entity.Category` dan `nil`.
  - method `StoreCategory` : menerima paramter `context.Context` dan `category` bertipe `Category`. Fungsi ini akan membuat data kategori baru berdasarkan `category` yang diberikan.
    - Jika terjadi error saat membuat data kategori, maka fungsi ini akan mengembalikan `0` dan `error`.
    - Jika kategori berhasil dibuat, maka fungsi ini akan mengembalikan hasil `id` category yang sudah dibuat dan `nil`.
  - method `StoreManyCategory` : menerima paramter `context.Context` dan `categories` bertipe `[]Category`. Fungsi ini akan membuat data kategori baru berdasarkan `categories` yang diberikan.
    - Jika terjadi error saat membuat data kategori, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika kategori berhasil dibuat, maka fungsi ini akan mengembalikan hasil `[]entity.Category` yang sudah dibuat dan `nil`.
  - method `UpdateCategory` : menerima paramter `context.Context` dan `category` bertipe `Category`. Fungsi ini akan mengupdate data kategori berdasarkan `category` yang diberikan.
    - Jika terjadi error saat mengupdate data kategori, maka fungsi ini akan mengembalikan `error`.
    - Jika kategori berhasil diupdate, maka fungsi ini akan mengembalikan `nil`.
  - method `GetCategoryByID`: menerima paramter `context.Context` dan `id` bertipe `int`. Fungsi ini akan mengembalikan data kategori berdasarkan `id` yang diberikan.
    - Jika terjadi error saat mengambil data kategori, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika kategori tidak ditemukan, maka fungsi ini akan mengembalikan `entity.Category{}` dan `nil`.
    - Jika kategori ditemukan, maka fungsi ini akan mengembalikan `entity.Category` dan `nil`.
  - method `DeleteCategory`: menerima paramter `context.Context` dan `id` bertipe `int`. Fungsi ini akan menghapus data kategori berdasarkan `id` yang diberikan.
    - Jika terjadi error saat menghapus data kategori, maka fungsi ini akan mengembalikan `error`.
    - Jika kategori berhasil dihapus, maka fungsi ini akan mengembalikan `nil`.

- **task**
  - method `GetTasks`: menerima paramter `context.Context` dan `id` bertipe `int` (_user id_). Fungsi ini akan mengembalikan data task berdasarkan user id di parameter `id` yang diberikan.
    - Jika terjadi error saat mengambil data task, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika task tidak ditemukan, maka fungsi ini akan mengembalikan `[]entity.Task{}` dan `nil`.
    - Jika task ditemukan, maka fungsi ini akan mengembalikan `[]entity.Task` dan `nil`.
  - method `StoreTask` : menerima paramter `context.Context` dan `task` bertipe `Task`. Fungsi ini akan membuat data task baru berdasarkan `task` yang diberikan.
    - Jika terjadi error saat membuat data task, maka fungsi ini akan mengembalikan `0` dan `error`.
    - Jika task berhasil dibuat, maka fungsi ini akan mengembalikan hasil `id` task yang sudah dibuat dan `nil`.
  - method `GetTaskByID` : menerima paramter `context.Context` dan `id` bertipe `int`. Fungsi ini akan mengembalikan data task berdasarkan `id` yang diberikan.
    - Jika terjadi error saat mengambil data task, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika task tidak ditemukan, maka fungsi ini akan mengembalikan `entity.Task{}` dan `nil`.
    - Jika task ditemukan, maka fungsi ini akan mengembalikan `entity.Task` dan `nil`.
  - method `GetTasksByCategoryID` : menerima paramter `context.Context` dan `id` bertipe `int` (_category id_). Fungsi ini akan mengembalikan data task berdasarkan category id di parameter `id` yang diberikan.
    - Jika terjadi error saat mengambil data task, maka fungsi ini akan mengembalikan `nil` dan `error`.
    - Jika task tidak ditemukan, maka fungsi ini akan mengembalikan `[]entity.Task{}` dan `nil`.
    - Jika task ditemukan, maka fungsi ini akan mengembalikan `[]entity.Task` dan `nil`.
  - method `UpdateTask` : menerima paramter `context.Context` dan `task` bertipe `Task`. Fungsi ini akan mengupdate data task berdasarkan `task` yang diberikan.
    - Jika terjadi error saat mengupdate data task, maka fungsi ini akan mengembalikan `error`.
    - Jika task berhasil diupdate, maka fungsi ini akan mengembalikan `nil`.
  - method `DeleteTask` : menerima paramter `context.Context` dan `id` bertipe `int`. Fungsi ini akan menghapus data task berdasarkan `id` yang diberikan.
    - Jika terjadi error saat menghapus data task, maka fungsi ini akan mengembalikan `error`.
    - Jika task berhasil dihapus, maka fungsi ini akan mengembalikan `nil`.

📁 **handler/api**

> Warning : abaikan code yang tidak berhubungan dengan instruksi di bawah ini pada folder `handler/api`

- **user**
  - method `Register` : adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`. Method ini akan membuat user baru dengan mememanggil userService.Register dengan parameter context dan *entity.User yang sudah didapatkan dari body request.
    - Method ini wajib mengirim data json dengan contoh format sebagai berikut:

      ```json
      {
        "fullname": <string>,
        "email": <string>,
        "password": <string>
      }
      ```

    - Jika data `fullname` atau `email` atau `password` kosong maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "register data is empty"
      }
      ```

    - Jika data `email` sudah terdaftar maka method ini akan mengembalikan response dengan status code `500` dan pesan error (bebaskan pesan errornya).

      ```json
      {
        "error": <error-message>
      }
      ```

    - Jika terjadi error saat menggunakan `userService.Register`, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
      {
        "error": "error internal server"
      }
      ```

    - Jika user berhasil dibuat, maka method ini akan mengembalikan response dengan status code `201` dan data user yang sudah dibuat sebagai berikut:

      ```json
      {
        "user_id": <int>,
        "message": "register success"
      }
      ```

  - method `Login`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`. Method ini akan melakukan login user dengan mememanggil userService.Login dengan parameter context dan *entity.User yang sudah didapatkan dari body request.
    - Method ini wajib mengirim data json dengan contoh format sebagai berikut:

      ```json
      {
        "email": <string>,
        "password": <string>
      }
      ```

    - Jika data `email` atau `password` kosong maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "email or password is empty"
      }
      ```

    - Jika terjadi error saat menggunakan `userService.Login`, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
      {
        "error": "error internal server"
      }
      ```

    - Jika user berhasil login, maka method ini akan mengembalikan response dengan status code `200` dan data user yang sudah login sebagai berikut:

      ```json
      {
        "user_id": <int>,
        "message": "login success"
      }
      ```

    - Jangan lupa untuk mengatur `cookie` menggunakan `http.SetCookie` dengan nama `user_id` dan value `user_id` yang sudah didapatkan dari `userService.Login`.

  - method `Logout`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`. Method ini akan melakukan logout user.
    - Method ini akan menghapus cookies yang sudah ada, jika berhasil akan mengembalika response dengan status code `200` dan pesan sebagai berikut:

      ```json
      {
        "message": "logout success"
      }
      ```

- **category**
  - method `GetCategory`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini akan mengambil data Context `id` yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Jika data context `id` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "invalid user id"
      }
      ```

    - Method ini akan memanggil `categoryService.GetCategory` dengan parameter `id` yang sudah didapatkan dari context. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
      {
        "error": "error internal server"
      }
      ```

    - Jika berhasil mengambil data category, maka method ini akan mengembalikan response dengan status code `200` dan data _list_ category bertipe `[]entity.Categiory`

  - method `CreateNewCategory`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini wajib mengirim data json dengan contoh format sebagai berikut:

      ```json
      {
        "type": "string"
      }
      ```

    - Jika data `type` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

        ```json
        {
            "error": "invalid category request"
        }
        ```

    - Method ini akan mengambil data Context `id` yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Jika data context `id` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "invalid user id"
      }
      ```

    - Method ini akan memanggil `categoryService.StoreCategory` dengan parameter `context.Context` dan data `entity.Category`. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
      {
        "error": "error internal server"
      }
      ```

    - Jika berhasil menyimpan data category, maka method ini akan mengembalikan response dengan status code `201` dan data category yang sudah disimpan.

    ```json
    {
        "user_id": <int>,
        "category_id": <int>,
        "message": "success create new category"
    }
    ```

  - method `DeleteCategory`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini akan mengambil data Context `id` yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Method ini akan mengambil data query URL `category_id` yang ada dari `http.Request`.

    ```go
    categoryID := r.URL.Query().Get("category_id")
    ```

    - Method ini akan memanggil `categoryService.DeleteCategory` dengan parameter `context.Context` dan `id` category yang didapat dari _query_ URL. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
        {
            "error": "error internal server"
        }
      ```

    - Jika berhasil menghapus data category, maka method ini akan mengembalikan response dengan status code `200` dan pesan sebagai berikut:

      ```json
        {
            "user_id": <int>,
            "category_id" : <int>,
            "message": "success delete category"
        }
      ```

- **task**
  - method `CreateNewTask`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini wajib mengirim data json dengan contoh format sebagai berikut:

      ```json
      {
        "title": <string>,
        "description": <string>,
        "category_id": <int>
      }
      ```

    - Jika data `title` atau `desripion` atau `category_id` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
        {
            "error": "invalid task request"
        }
       ```

    - Method ini akan mengambil data Context `id` yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Jika data context `id` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "invalid user id"
      }
      ```

    - Method ini akan memanggil `taskService.StoreTask` dengan parameter `context.Context` dan data `entity.Task`. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
      {
        "error": "error internal server"
      }
      ```

    - Jika berhasil menyimpan data task, maka method ini akan mengembalikan response dengan status code `201` dan data task yang sudah disimpan.

      ```json
      {
            "user_id": <int>,
            "task_id": <int>,
            "message": "success create new task"
      }
      ```

  - method `GetTask`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini akan mengambil data Context `id` yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Jika data context `id` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "invalid user id"
      }
      ```

    - Method ini akan mendapatkan query URL `task_id` dari `http.Request`.
      - Jika `task_id` isinya kosong, maka method ini akan memanggil `taskService.GetTasks` dengan parameter `context.Context` dan `id` user yang sudah didapatkan dari context. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

        ```json
        {
          "error": "error internal server"
        }
        ```

      - Jika berhasil mengambil data task, maka method ini akan mengembalikan response dengan status code `200` dan data _list_ task bertipe `[]entity.Task`

      - Jika `task_id` isinya tidak kosong, maka method ini akan memanggil `taskService.GetTaskByID` dengan parameter `context.Context` dan `id` task yang sudah didapatkan dari query URL. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

        ```json
        {
          "error": "error internal server"
        }
        ```

      - Jika berhasil mengambil data task, maka method ini akan mengembalikan response dengan status code `200` dan data task bertipe `entity.Task`

  - method `UpdateTask`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini wajib mengirim data json dengan contoh format sebagai berikut:

      ```json
      {
        "title": <string>,
        "description": <string>,
        "category_id": <int>
      }
      ```

    - Method ini akan mengambil data Context `id` (user_id) yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Jika data context `id` kosong, maka method ini akan mengembalikan response dengan status code `400` dan pesan error sebagai berikut:

      ```json
      {
        "error": "invalid user id"
      }
      ```

    - Method ini akan memanggil `taskService.UpdateTask` dengan parameter `context.Context`dan data `entity.Task`. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
      {
        "error": "error internal server"
      }
      ```

    - Jika berhasil mengupdate data task, maka method ini akan mengembalikan response dengan status code `200` dan data task yang sudah diupdate.

      ```json
      {
            "user_id": <int>,
            "task_id": <int>,
            "message": "success update task"
      }
      ```

  - method `DeleteTask`: adalah sebuah handler yang menerima parameter `http.ResponseWriter` dan `*http.Request`.
    - Method ini akan mengambil data Context `id` yang sudah di set sebelumnya menggunakan `middleware.Auth`.

    - Method ini akan mengambil data query URL `task_id` yang ada dari `http.Request`.

    ```go
    taskID := r.URL.Query().Get("task_id")
    ```

    - Method ini akan memanggil `taskService.DeleteTask` dengan parameter `context.Context` dan `id` task yang didapat dari _query_ URL. Jika terjadi error, maka method ini akan mengembalikan response dengan status code `500` dan pesan error sebagai berikut:

      ```json
        {
            "error": "error internal server"
        }
      ```

    - Jika berhasil menghapus data task, maka method ini akan mengembalikan response dengan status code `200` dan pesan sebagai berikut:

      ```json
        {
            "user_id": <int>,
            "task_id" : <int>,
            "message": "success delete task"
        }
      ```
