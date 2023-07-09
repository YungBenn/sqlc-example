package todo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"sqlc-example/pkg/sqlc"
)

type Repository interface {
	sqlc.Querier
}

type SQLRepository struct {
	db *pgxpool.Pool
	*sqlc.Queries
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &SQLRepository{
		db:      db,
		Queries: sqlc.New(db),
	}
}
