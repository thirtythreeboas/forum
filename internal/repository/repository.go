package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

type Interface interface {
	GetPosts(ctx context.Context) string
}

func New(db *pgxpool.Pool) Interface {
	return &repository{db: db}
}
