package repository

import (
	"context"
	"fmt"
	"forum/internal/model"

	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateForum(ctx context.Context, newForum *model.ForumCreate) (*model.Forum, error) {
	const op = "repository.forum.CreateForum"

	query := `
		SELECT id, "user", title, slug, posts, threads
		FROM forums
		WHERE "user" = $1
	`

	var forum model.Forum

	err := r.db.QueryRow(ctx, query, newForum.User).Scan(
		&forum.ID,
		&forum.Posts,
		&forum.Slug,
		&forum.Threads,
		&forum.Title,
		&forum.User,
	)

	if err == nil {
		return &forum, nil
	}

	if err != pgx.ErrNoRows {
		return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
	}

	queryInsert := `
		INSERT INTO forums ("user", title, slug)
		VALUES ($1, $2, $3)
		RETURNING id, "user", title, slug, posts, threads
	`

	err = r.db.QueryRow(ctx, queryInsert, newForum.User, newForum.Title, newForum.Slug).Scan(
		&forum.ID,
		&forum.Posts,
		&forum.Slug,
		&forum.Threads,
		&forum.Title,
		&forum.User,
	)

	if err != nil {
		return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
	}

	return &forum, nil
}
