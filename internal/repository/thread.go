package repository

import (
	"context"
	"errors"
	"forum/internal/model"
	cerr "forum/pkg/errors"

	"github.com/jackc/pgx/v5"
)

func (r *repository) GetThreadInfo(ctx context.Context, forumSlug string) (*model.Thread, error) {
	query := `
		SELECT (title, author, forum, message, votes, slug, created)
		FROM threads
		WHERE slug = $1
	`

	var thread model.Thread

	err := r.db.QueryRow(ctx, query, forumSlug).Scan(
		&thread.Title,
		&thread.Author,
		&thread.Forum,
		&thread.Message,
		&thread.Votes,
		&thread.Slug,
		&thread.Created,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, cerr.ErrThreadDoesntExist
		}
		return nil, err
	}

	return &thread, nil
}
