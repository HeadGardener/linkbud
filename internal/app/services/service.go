package services

import "github.com/HeadGardener/linkbud/internal/app/repository"

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
