package repository

import (
	"context"
	"errors"
	"fmt"
	"forum/internal/model"
	cerr "forum/pkg/errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) CreateForum(ctx context.Context, newForum *model.ForumCreate) (*model.Forum, error) {
	const op = "repository.forum.CreateForum"

	query := `
		SELECT title, "user", slug, posts, threads
		FROM forums
		WHERE "user" = $1
	`

	var forum model.Forum

	err := r.db.QueryRow(ctx, query, newForum.User).Scan(
		&forum.Title,
		&forum.User,
		&forum.Slug,
		&forum.Posts,
		&forum.Threads,
	)

	if err == nil {
		return &forum, cerr.ErrForumAlreadyExists
	}

	queryInsert := `
		INSERT INTO forums ("user", title, slug)
		VALUES ($1, $2, $3)
		RETURNING  title, "user", slug, posts, threads
	`

	err = r.db.QueryRow(ctx, queryInsert, newForum.User, newForum.Title, newForum.Slug).Scan(
		&forum.Title,
		&forum.User,
		&forum.Slug,
		&forum.Posts,
		&forum.Threads,
	)

	if err != nil {
		return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
	}

	return &forum, nil
}

func (r *repository) GetForum(ctx context.Context, slug string) (*model.Forum, error) {
	query := `
		SEELCT (title, "user", slug, posts, threads)
		FROM forums
		WHERE slug = $1
	`

	forum := &model.Forum{}

	err := r.db.QueryRow(ctx, query, slug).Scan(
		&forum.Title,
		&forum.User,
		&forum.Slug,
		&forum.Posts,
		&forum.Threads,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, cerr.ErrForumDoesntExist
		}
		return nil, err
	}

	return forum, nil
}

func (r *repository) CreateThread(ctx context.Context, newThread *model.NewThread) (*model.Thread, error) {
	const op = "repository.forum.CreateThread"

	threadSlug := strings.ToLower(newThread.Title)
	threadSlug = strings.ReplaceAll(threadSlug, " ", "-")

	query := `
		INSERT INTO threads (title, author, forum, message, slug)
		VALUES ($1, $2, 3$, $4, $5)
		RETURNING title, author, forum, message, votes, slug, created
	`

	var thread model.Thread

	err := r.db.QueryRow(ctx, query, newThread.Title, newThread.Author, newThread.Slug, newThread.Message, threadSlug).Scan(
		&thread.Title,
		&thread.Author,
		&thread.Forum,
		&thread.Message,
		&thread.Votes,
		&thread.Slug,
		&thread.Created,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, cerr.ErrThreadAlreadyExists
			}
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &thread, nil
}
