package services

import (
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/HeadGardener/linkbud/internal/app/repository"
	authservice "github.com/HeadGardener/linkbud/internal/app/services/auth"
	linkservice "github.com/HeadGardener/linkbud/internal/app/services/link"
	listservice "github.com/HeadGardener/linkbud/internal/app/services/list"
)

type Authorization interface {
	Create(user models.User) (int, error)
	GenerateToken(userInput models.UserInput) (string, error)
	ParseToken(token string) (int, error)
	CheckUsername(username string) (int, error)
}

type ListInterface interface {
	Create(userID int, listInput models.ListInput) (int, error)
	GetAll(userID int) ([]models.List, error)
	GetList(userID int, title string) (models.List, error)
	Update(userID int, title string, listInput models.ListInput) (int, error)
	Delete(userID int, title string) (int, error)
}

type LinkInterface interface {
	Create(userID int, link models.Link, listID int, title string) (int, error)
	GetAll(userID int, listID int, title string) ([]models.Link, error)
	GetByID(userID int, listID int, title string, linkTitle string) (models.Link, error)
	Update(userID, listID int, linkInput models.LinkInput, title string) error
	Delete(userID, linkID int, title string) error
}

type Service struct {
	Authorization
	ListInterface
	LinkInterface
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: authservice.NewAuthService(repository),
		ListInterface: listservice.NewListService(repository),
		LinkInterface: linkservice.NewLinkService(repository),
	}
}
