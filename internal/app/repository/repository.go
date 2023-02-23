package repository

import (
	"github.com/HeadGardener/linkbud/internal/app/models"
	authrepository "github.com/HeadGardener/linkbud/internal/app/repository/auth"
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
}

type Repository struct {
	Authorization
	ListInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: authrepository.NewAuthRepository(db),
		ListInterface: listrepository.NewListRepository(db),
	}
}
