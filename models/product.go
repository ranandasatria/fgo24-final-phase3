package models

import (
	"context"
	"test-fase-3/utils"
	"time"

	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID         int       `json:"id"`
	Nama       string    `json:"nama"`
	GambarURL  string    `json:"gambar_url"`
	KategoriID int       `json:"kategori_id"`
	Stok       int       `json:"stok"`
	HargaBeli  int       `json:"harga_beli"`
	HargaJual  int       `json:"harga_jual"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductWithCategory struct {
	ID            int       `json:"id"`
	Nama          string    `json:"nama"`
	GambarURL     string    `json:"gambar_url"`
	KategoriID    int       `json:"kategori_id"`
	KategoriNama  string    `json:"kategori_nama"`
	Stok          int       `json:"stok"`
	HargaBeli     int       `json:"harga_beli"`
	HargaJual     int       `json:"harga_jual"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func CreateProduct(p Product) error {
	db, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
		INSERT INTO produk (nama, gambar_url, kategori_id, stok, harga_beli, harga_jual, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = db.Exec(context.Background(), query,
		p.Nama, p.GambarURL, p.KategoriID, p.Stok, p.HargaBeli, p.HargaJual, time.Now(), time.Now())
	return err
}

func GetAllProducts() ([]ProductWithCategory, error) {
	db, err := utils.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT 
			p.id, p.nama, p.gambar_url, p.kategori_id, k.nama AS kategori_nama,
			p.stok, p.harga_beli, p.harga_jual, p.created_at, p.updated_at
		FROM produk p
		JOIN kategori_produk k ON p.kategori_id = k.id
	`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[ProductWithCategory])
	return products, err
}