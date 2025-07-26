package postgres

import (
	"context"
	"fmt"
	"forum/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MustNew(ctx context.Context, cfg *config.PGConfig) *pgxpool.Pool {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.Mode,
	)

	fmt.Printf("conn: %s\n", conn)

	dbPool, err := pgxpool.New(ctx, conn)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	return dbPool
}
