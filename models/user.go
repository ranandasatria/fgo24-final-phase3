package models

import (
	"context"
	"test-fase-3/utils"
	"time"
)

type User struct {
	ID        int
	Nama      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user User) error {
	db, err := utils.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	INSERT INTO users(nama, email, password, role, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6)
	`

	_, err = db.Exec(context.Background(), query,
			user.Nama, user.Email, user.Password, user.Role, time.Now(), time.Now(),
	)

	return err
}

func FindUserByEmail(email string) (User, error) {
	db, err := utils.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	query := `SELECT id, nama, email, password, role, created_at, updated_at FROM users WHERE email = $1 LIMIT 1`

	var user User
	err = db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Nama,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
