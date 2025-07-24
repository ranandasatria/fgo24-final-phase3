package main

import (
	"context"
	"fmt"
	"time"

	"test-fase-3/utils"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	err := Seed()
	if err != nil {
		fmt.Println("Seeding gagal:", err)
		return
	}
	fmt.Println("Seeding sukses.")
}

func Seed() error {
	db, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	ctx := context.Background()

	password := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, `
		INSERT INTO users (nama, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (email) DO NOTHING;
	`, "Admin", "admin@mail.com", string(hashedPassword), "admin", time.Now(), time.Now())
	if err != nil {
		return err
	}
	fmt.Println("Seeded admin user.")

	kategoris := []struct {
		nama, deskripsi string
	}{
		{"Elektronik", "Produk-produk elektronik seperti TV, laptop, dll"},
		{"Pakaian", "Berbagai jenis pakaian pria dan wanita"},
		{"Makanan", "Produk makanan dan minuman"},
	}

	for _, k := range kategoris {
		_, err := db.Exec(ctx, `
			INSERT INTO kategori_produk (nama, deskripsi, created_at, updated_at)
			VALUES ($1, $2, $3, $4);
		`, k.nama, k.deskripsi, time.Now(), time.Now())
		if err != nil {
			return err
		}
	}

	fmt.Println("Seeded kategori_produk.")
	return nil
}
