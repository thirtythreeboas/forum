package repository

import (
	"context"
	"errors"
	"fmt"
	"forum/internal/model"
	cerr "forum/pkg/errors"

	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateUser(ctx context.Context, user *model.NewProfile) (*model.User, error) {
	const op = "repository.user.CreateUser"

	query := `
		INSERT INTO users (nickname, fullname, email, about)
		VALUES ($1, $2, $3, $4)
		RETURNING id, nickname, fullname, email, about
	`

	var newUser model.User

	err := r.db.QueryRow(ctx, query, user.Nickname, user.Fullname, user.Email, user.About).Scan(
		&newUser.Id,
		&newUser.Nickname,
		&newUser.Fullname,
		&newUser.Email,
		&newUser.About,
	)

	if err != nil {
		return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
	}

	return &newUser, nil
}

func (r *repository) GetProfile(ctx context.Context, nickname string) (*model.User, error) {
	const op = "repository.user.GetProfile"

	query := `
		SELECT nickname, fullname, email, about
		FROM users
		WHERE nickname = $1
	`
	var existingUser model.User

	err := r.db.QueryRow(ctx, query, nickname).Scan(
		&existingUser.Nickname,
		&existingUser.Fullname,
		&existingUser.Email,
		&existingUser.About,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, cerr.ErrUserDoesntExist
		}

		return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
	}

	return &existingUser, cerr.ErrUserAlreadyExists
}

func (r *repository) ChangeProfile(ctx context.Context, user *model.User) (*model.User, error) {
	const op = "repository.user.ChangeProfile"

	var existingNickname string
	checkEmailQuery := `
			SELECT nickname
			FROM users
			WHERE email = $1 AND nickname != $2
	`

	err := r.db.QueryRow(ctx, checkEmailQuery, user.Email, user.Nickname).Scan(&existingNickname)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
		}
	}

	if err == nil {
		return nil, cerr.ErrEmailIsInUse
	}

	updateQuery := `
		UPDATE users
		SET
			fullname = $1,
			about = $2,
			email = $3
		WHERE nickname = $4
		RETURNING nickname, fullname, about, email
	`

	var updatedUser model.User

	err = r.db.QueryRow(ctx, updateQuery, user.Fullname, user.About, user.Email, user.Nickname).Scan(
		&updatedUser.Nickname,
		&updatedUser.Fullname,
		&updatedUser.About,
		&updatedUser.Email,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, cerr.ErrUserDoesntExist
		}
		return nil, fmt.Errorf("location: %s | error: %v", op, err.Error())
	}

	return &updatedUser, nil
}
