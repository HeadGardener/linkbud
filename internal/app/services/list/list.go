package listservice

import (
	"errors"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/HeadGardener/linkbud/internal/app/repository"
	"strings"
)

type ListService struct {
	repos *repository.Repository
}

func NewListService(repos *repository.Repository) *ListService {
	return &ListService{repos: repos}
}

func (s *ListService) Create(userID int, list models.LinkList) (int, error) {
	if err := validateList(&list); err != nil {
		return 0, err
	}

	return s.repos.ListInterface.Create(userID, list)
}

func validateList(list *models.LinkList) error {
	if len(list.Title) > 25 {
		return errors.New("list title is too long")
	}

	list.ShortTitle = strings.Join(strings.Split(strings.TrimSpace(list.Title), " "), "")

	return nil
}
