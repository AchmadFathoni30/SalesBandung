# Sales Bandung API

## Technology
- **Go 1.23.2**
- **Gin Gonic**
- **GORM**
- **PostgreSQL**

## Fitur
- **Basic Authentication**: Menggunakan basic authentication untuk akses endpoint API.
- **Database**: Koneksi ke PostgreSQL menggunakan GORM dengan konfigurasi dari file `.env`.
- **API Endpoint**: Mengambil data dari tabel :
                    `sales_hdr_bdg`, 
                    `sales_dtl_bdg`, 
                    `sales_payment_bdg`, 
                    `sales_void_bdg`, 
                    `sales_dp_bdg`
- **Docker**: Menyediakan Dockerfile untuk membangun aplikasi dalam container Docker.

## Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda memiliki perangkat lunak berikut terinstal:
- **Go 1.23.2** 
- **Docker** 

### 2. **Instalasi dependency dengan Go mod**

```bash
go mod tidy