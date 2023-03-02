package repository

import (
	"github.com/HeadGardener/linkbud/internal/app/models"
	authrepository "github.com/HeadGardener/linkbud/internal/app/repository/auth"
	linkrepository "github.com/HeadGardener/linkbud/internal/app/repository/link"
	listrepository "github.com/HeadGardener/linkbud/internal/app/repository/list"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user models.User) (int, error)
	IfUserExist(userInput models.UserInput) (int, error)
	IfUserExistByUN(username string) (int, error)
}

type ListInterface interface {
	Create(userID int, list models.List) (int, error)
	GetAll(userID int) ([]models.List, error)
	GetList(userID int, title string) (models.List, error)
	Update(userID int, title string, list models.ListInput) (int, error)
	Delete(userID int, title string) (int, error)
	CheckIfTitleExists(userID int, title string) error
}

type LinkInterface interface {
	Create(link models.Link, listID int) (int, error)
	GetAll(listID int) ([]models.Link, error)
	GetByID(listID int, linkTitle string) (models.Link, error)
	Update(listID int, linkInput models.LinkInput) error
	Delete(linkID int) error
}

type Repository struct {
	Authorization
	ListInterface
	LinkInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: authrepository.NewAuthRepository(db),
		ListInterface: listrepository.NewListRepository(db),
		LinkInterface: linkrepository.NewLinkRepository(db),
	}
}
