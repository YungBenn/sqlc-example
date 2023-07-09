package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func Connect(ctx context.Context, db_url string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, db_url)

	if err != nil {
		return nil, err
	}

	fmt.Println("Success connect to database")

	return db, nil
}
