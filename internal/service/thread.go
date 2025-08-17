package service

import (
	"context"
	"forum/internal/model"
)

func (s *Service) GetThreadInfo(slug string) (*model.Thread, error) {
	thread, err := s.repo.GetThreadInfo(context.Background(), slug)

	if err != nil {
		return nil, err
	}

	return thread, err
}
