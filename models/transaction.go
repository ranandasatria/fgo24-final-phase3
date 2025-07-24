package models

import (
	"context"
	"errors"
	"fmt"
	"test-fase-3/utils"
	"time"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	ID             int
	ProdukID       int
	UserID         int
	Tipe           string
	Jumlah         int
	TotalHargaBeli int
	TotalHargaJual int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TransactionWithStock struct {
	ID                  int    `json:"id"`
	NamaProduk          string `json:"namaProduk"`
	NamaKategori        string `json:"namaKategori"`
	BarangMasuk         int    `json:"barangMasuk"`
	BarangKeluar        int    `json:"barangKeluar"`
	HargaBeli           int    `json:"hargaBeli"`
	HargaJual           int    `json:"hargaJual"`
	TotalHargaPembelian *int   `json:"totalHargaPembelian" db:"total_harga_beli"`
	TotalHargaPenjualan *int   `json:"totalHargaPenjualan" db:"total_harga_jual"`
	StokTersedia        int    `json:"stokTersedia"`
	CreatedAt           string `json:"createdAt"`
}

func CreateTransaction(t Transaction) error {
	db, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var (
		hargaBeli int
		hargaJual int
		stok      int
	)

	err = db.QueryRow(context.Background(), `
		SELECT harga_beli, harga_jual, stok FROM produk WHERE id = $1
	`, t.ProdukID).Scan(&hargaBeli, &hargaJual, &stok)
	if err != nil {
		return err
	}

	if t.Tipe == "OUT" && t.Jumlah > stok {
		return errors.New("jumlah melebihi stok yang tersedia")
	}

	t.TotalHargaBeli = hargaBeli * t.Jumlah
	t.TotalHargaJual = hargaJual * t.Jumlah

	_, err = db.Exec(context.Background(), `
		INSERT INTO transaksi (produk_id, user_id, tipe, jumlah, total_harga_beli, total_harga_jual)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, t.ProdukID, t.UserID, t.Tipe, t.Jumlah, t.TotalHargaBeli, t.TotalHargaJual)
	if err != nil {
		return err
	}

	var newStok int
	if t.Tipe == "IN" {
		newStok = stok + t.Jumlah
	} else {
		newStok = stok - t.Jumlah
	}

	_, err = db.Exec(context.Background(), `
		UPDATE produk SET stok = $1, updated_at = NOW() WHERE id = $2
	`, newStok, t.ProdukID)

	return err
}
func GetAllTransactions() ([]TransactionWithStock, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		fmt.Printf("ConnectDB error: %v\n", err)
		return nil, err
	}
	defer conn.Close()

	query := `
		SELECT 
			t.id,
			p.nama AS nama_produk,
			COALESCE(k.nama, '') AS nama_kategori,
			CASE WHEN t.tipe = 'IN' THEN t.jumlah ELSE 0 END AS barang_masuk,
			CASE WHEN t.tipe = 'OUT' THEN t.jumlah ELSE 0 END AS barang_keluar,
			p.harga_beli,
			p.harga_jual,
			t.total_harga_beli,
			t.total_harga_jual,
			(
				SELECT COALESCE(SUM(CASE WHEN tipe = 'IN' THEN jumlah ELSE -jumlah END), 0)
				FROM transaksi
				WHERE produk_id = t.produk_id
				AND created_at <= t.created_at
			) AS stok_tersedia,
			t.created_at::text
		FROM transaksi t
		JOIN produk p ON t.produk_id = p.id
		LEFT JOIN kategori_produk k ON p.kategori_id = k.id
		ORDER BY t.created_at DESC
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByName[TransactionWithStock])
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
