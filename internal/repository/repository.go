package repository

import (
	"context"
	"forum/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

type Interface interface {
	CreateForum(ctx context.Context, newForum *model.ForumCreate) (*model.Forum, error)
	CreateUser(ctx context.Context, user *model.NewProfile) (*model.User, error)
	GetProfile(ctx context.Context, nickname string) (*model.User, error)
	ChangeProfile(ctx context.Context, user *model.User) (*model.User, error)
}

func New(db *pgxpool.Pool) Interface {
	return &repository{db: db}
}
