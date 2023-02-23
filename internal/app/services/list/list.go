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

func (s *ListService) Create(userID int, listInput models.ListInput) (int, error) {
	list := listFromInput(listInput)

	if err := checkTitleLen(list.Title); err != nil {
		return 0, err
	}

	list.ShortTitle = makeShortTitle(list.Title)

	return s.repos.ListInterface.Create(userID, list)
}

func (s *ListService) GetAll(userID int) ([]models.List, error) {
	return s.repos.ListInterface.GetAll(userID)
}

func (s *ListService) GetList(userID int, title string) (models.List, error) {
	if err := checkTitleLen(title); err != nil {
		return models.List{}, err
	}

	return s.repos.ListInterface.GetList(userID, title)
}

func (s *ListService) Update(userID int, title string, listInput models.ListInput) (int, error) {
	if err := checkTitleLen(title); err != nil {
		return 0, err
	}

	if err := checkTitleLen(*listInput.Title); err != nil {
		return 0, err
	}

	listInput.ShortTitle = makeShortTitle(*listInput.Title)

	return s.repos.ListInterface.Update(userID, title, listInput)
}

func (s *ListService) Delete(userID int, title string) (int, error) {
	if err := checkTitleLen(title); err != nil {
		return 0, err
	}

	return s.repos.ListInterface.Delete(userID, title)
}

func makeShortTitle(title string) string {
	if len(title) > 25 {
		return ""
	}

	return strings.Join(strings.Split(strings.TrimSpace(title), " "), "")
}

func checkTitleLen(title string) error {
	if len(title) > 25 {
		return errors.New("list title is too long")
	}

	return nil
}

func listFromInput(l models.ListInput) models.List {
	var list models.List
	list.Title = *l.Title
	list.Description = *l.Description

	return list
}
