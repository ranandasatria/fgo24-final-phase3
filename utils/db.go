package utils

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	url := "postgres://postgres:1@localhost:5432/postgres"

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return pool, err
}
