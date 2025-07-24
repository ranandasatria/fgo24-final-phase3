## User Requirement

#### Authentication
- User dapat melakukan **register** dengan email dan password
- User dapat **login** dengan email dan password
- Jika login berhasil, sistem akan memberikan **JWT token**
- User hanya bisa mengakses fitur utama setelah login

#### Dashboard
- Setelah login, user diarahkan ke halaman **dashboard**
- Di dashboard, user bisa memilih:
  - **Melihat daftar produk**
  - **Melihat riwayat transaksi**

#### Produk
- User dapat melihat daftar semua produk
- User dengan role admin dapat menambahkan produk baru
- Setelah produk ditambahkan, daftar produk akan diperbarui

#### Transaksi
- User dapat melihat riwayat transaksi masuk/keluar
- Admin dapat menambahkan transaksi baru (masuk / keluar stok)
- Setelah transaksi ditambahkan, riwayat akan diperbarui



## Flowchart

```mermaid
flowchart TD
    %% AUTH FLOW
    A([Start])
    B[Register]
    C{Input Valid?}
    D[Create User]
    E[Login]
    F{Credential Valid?}
    G[Generate JWT Token]
    H[Dashboard]

    %% MAIN FEATURES (CABANG DARI DASHBOARD)
    I[View Product List]
    J[Add Product]
    K[View Transaction History]
    L[Stock In / Stock Out]
    M([End])

    %% FLOW CONNECTIONS
    A --> B --> C
    C -- No --> B
    C -- Yes --> D --> E
    E --> F
    F -- No --> E
    F -- Yes --> G --> H

    %% DASHBOARD CABANG
    H --> I
    H --> K

    %% PRODUCT FLOW
    I --> J --> I

    %% TRANSACTION FLOW
    K --> L --> K
    K --> M

```

### ERD
```mermaid
erDiagram
 user ||--o{ transaksi : "creates"
produk ||--o{ transaksi : "has"
kategori_produk ||--o{ produk : "contains"

    user {
        int id PK
        varchar nama
        varchar email
        varchar password
        varchar role
        timestamp created_at
        timestamp updated_at
    }

    kategori_produk {
        int id PK
        varchar nama
        text deskripsi
        timestamp created_at
        timestamp updated_at
    }

    produk {
        int id PK
        varchar nama
        varchar gambar_url
        int kategori_id FK
        int stok
        int harga_beli
        int harga_jual
        timestamp created_at
        timestamp updated_at
    }

    transaksi {
        int id PK
        int produk_id FK
        int user_id FK
        enum tipe 
        int jumlah
        int total_harga_beli
        int total_harga_jual
        timestamp created_at
        timestamp updated_at
    }

   

```