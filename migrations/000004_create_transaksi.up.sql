CREATE TYPE tipe_transaksi AS ENUM ('IN', 'OUT');

CREATE TABLE transaksi (
    id SERIAL PRIMARY KEY,
    produk_id INT REFERENCES produk(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    tipe tipe_transaksi NOT NULL,
    jumlah INT NOT NULL,
    total_harga_beli INT,
    total_harga_jual INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
