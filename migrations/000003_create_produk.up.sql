CREATE TABLE produk (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    gambar_url VARCHAR(255),
    kategori_id INT REFERENCES kategori_produk(id) ON DELETE SET NULL,
    stok INT NOT NULL DEFAULT 0,
    harga_beli INT NOT NULL,
    harga_jual INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
