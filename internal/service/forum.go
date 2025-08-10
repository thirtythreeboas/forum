package service

import (
	"context"
	"forum/internal/model"
)

func (s *Service) CreateForum(newForum *model.ForumCreate) (*model.Forum, error) {
	forum, err := s.repo.CreateForum(context.Background(), newForum)

	if err != nil {
		return nil, err
	}

	return forum, nil
}
