package service

import (
	"forum/internal/repository"
)

type Service struct {
	repo repository.Interface
}

func NewService(repo repository.Interface) *Service {
	return &Service{
		repo: repo,
	}
}
