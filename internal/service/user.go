package service

import (
	"context"
	"forum/internal/model"
	cerr "forum/pkg/errors"
)

func (s *Service) CreateUser(newUser *model.NewProfile) (*model.User, error) {
	existingUser, err := s.repo.GetProfile(context.Background(), newUser.Nickname)

	if err != nil {
		if err == cerr.ErrUserAlreadyExists {
			return existingUser, err
		}
		if err != cerr.ErrUserDoesntExist {
			return nil, err
		}
	}

	user, err := s.repo.CreateUser(context.Background(), newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUser(nickname string) (*model.User, error) {
	user, err := s.repo.GetProfile(context.Background(), nickname)

	if err != nil {
		if err == cerr.ErrUserDoesntExist {
			return nil, cerr.ErrCantFindUser
		}

		if err != cerr.ErrUserAlreadyExists {
			return nil, err
		}
	}

	return user, nil
}

func (s *Service) ChangeProfile(user *model.User) (*model.User, error) {
	extractedUser, err := s.repo.ChangeProfile(context.Background(), user)

	if err != nil {
		if err == cerr.ErrEmailIsInUse {
			return nil, cerr.ErrEmailIsInUse
		}
		if err == cerr.ErrUserDoesntExist {
			return nil, cerr.ErrUserDoesntExist
		}
		return nil, err
	}

	return extractedUser, nil
}
