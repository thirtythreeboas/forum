package service

import (
	"context"
	"forum/internal/model"
	cerr "forum/pkg/errors"
)

func (s *Service) CreateForum(newForum *model.ForumCreate) (*model.Forum, error) {
	_, err := s.repo.GetProfile(context.Background(), newForum.User)
	if err != nil {
		if err == cerr.ErrUserDoesntExist {
			return nil, cerr.ErrUserDoesntExist
		}
		if err != cerr.ErrUserAlreadyExists {
			return nil, err
		}
	}

	forum, err := s.repo.CreateForum(context.Background(), newForum)
	if err != nil {
		if err == cerr.ErrForumAlreadyExists {
			return forum, err
		}
		return nil, err
	}

	return forum, nil
}

func (s *Service) GetForumInfo(slug string) (*model.Forum, error) {
	forum, err := s.repo.GetForum(context.Background(), slug)
	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (s *Service) CreateThread(newThread *model.NewThread) (*model.Thread, error) {
	_, err := s.repo.GetProfile(context.Background(), newThread.Author)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.GetForum(context.Background(), newThread.Slug)
	if err != nil {
		return nil, err
	}

	thread, err := s.repo.CreateThread(context.Background(), newThread)
	if err != nil {
		if err == cerr.ErrThreadAlreadyExists {
			existingThread, err := s.repo.GetThreadInfo(context.Background(), newThread.Slug)
			if err != nil {
				if err != cerr.ErrThreadDoesntExist {
					return nil, err
				}
				return nil, err
			}

			return existingThread, cerr.ErrThreadAlreadyExists
		}
		return nil, err
	}

	return thread, nil
}
