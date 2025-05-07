# API PRODUCT with JWT Auth

Product API adalah RESTful API sederhana untuk mengelola data produk.
Dibangun menggunakan Golang, Echo framework, GORM dan MySQL sebagai database.

## Tools

- ![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
- ![Mysql Database](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white)
- ![Insomnia](https://img.shields.io/badge/Insomnia-5849be?style=for-the-badge&logo=Insomnia&logoColor=white) ![](https://img.shields.io/badge/for%20API%20Testing-8A2BE2)

## Installation

#### 1. Clone Repository

```bash
$   git clone https://product_api
$   cd product_api
```

#### 2. Setup Database

- Pastikan MySql sudah terinstall
- Buat database baru, misal `product_api`
- Buat file `.env` sesuaikan `DB_HOST` `DB_PORT` `DB_USERNAME` `DB_PASSWORD`
- Sesuaikan konfigurasi koneksi database di file `config/database.go `

`.env `

```bash
    LOCALHOST=127.0.0.1
    APP_PORT=8080

    DB_DATABASE=<contoh_nama_database>
    DB_HOST=<contoh_host>
    DB_PORT=<contoh_port>
    DB_USERNAME=<contoh_username>
    DB_PASSWORD=<contoh_password>

    //Untuk Generate Token
    JWT_SECRET=<disesuikan_kembali>
```

`config/database.go `

```bash
    db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
```

#### 3. Install Dependency

```bash
$   go mod tidy
```

#### 4. Run Aplikasi

```bash
$   go run main.go
```

Aplikasi akan berjalan di `127.0.0.1:8080`

## API Request

#### Registrasi

```http
  POST /user/register
```

| Parameter  | Type     | Description                        |
| :--------- | :------- | :--------------------------------- |
| `username` | `string` | **Required**. Your Username        |
| `role`     | `string` | **Required**. Your Role for Access |
| `email`    | `string` | **Required**. Your Email           |
| `password` | `string` | **Required**. Your Password        |

#### Login

```http
  POST /user/login
```

| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `username` | `string` | **Required**. Your Username |
| `password` | `string` | **Required**. Your Password |

#### Create Product

![](https://img.shields.io/badge/require%20JWT%20for%20Authentication-3273a8)

```http
  POST /product
```

| Parameter      | Type      | Description                       |
| :------------- | :-------- | :-------------------------------- |
| `product_name` | `string`  | **Required**. Name of The Product |
| `total`        | `integer` | Amount of The Product             |
| `price`        | `float`   | Price of The Product              |

#### Get Product by User ID

![](https://img.shields.io/badge/require%20JWT%20for%20Authentication-3273a8)

```http
  GET /product
```

User ID akan didapatkan dari Token yang sudah di Decode

#### Update Product

![](https://img.shields.io/badge/require%20JWT%20for%20Authentication-3273a8)

```http
  PATCH /product/:product_id
```

| Parameter      | Type      | Description                       |
| :------------- | :-------- | :-------------------------------- |
| `product_name` | `string`  | **Required**. Name of The Product |
| `total`        | `integer` | Amount of The Product             |
| `price`        | `float`   | Price of The Product              |

#### Delete by User ID

![](https://img.shields.io/badge/require%20JWT%20for%20Authentication-3273a8)

```http
  DELETE /product/:product_id
```

User ID akan didapatkan dari Token yang sudah di Decode
