package linkservice

import (
	"errors"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/HeadGardener/linkbud/internal/app/repository"
)

type LinkService struct {
	repos *repository.Repository
}

func NewLinkService(repos *repository.Repository) *LinkService {
	return &LinkService{repos: repos}
}

func (s *LinkService) Create(userID int, link models.Link, listID int, title string) (int, error) {
	if err := s.repos.ListInterface.CheckIfTitleExists(userID, title); err == nil {
		return 0, errors.New("title don't exist")
	}

	return s.repos.LinkInterface.Create(link, listID)
}

func (s *LinkService) GetAll(userID int, listID int, title string) ([]models.Link, error) {
	if err := s.repos.ListInterface.CheckIfTitleExists(userID, title); err == nil {
		return nil, errors.New("title don't exist")
	}

	return s.repos.LinkInterface.GetAll(listID)
}

func (s *LinkService) GetByID(userID int, listID int, title string, linkTitle string) (models.Link, error) {
	if err := s.repos.ListInterface.CheckIfTitleExists(userID, title); err == nil {
		return models.Link{}, errors.New("title don't exist")
	}

	return s.repos.LinkInterface.GetByID(listID, linkTitle)
}

func (s *LinkService) Update(userID, listID int, linkInput models.LinkInput, title string) error {
	if err := s.repos.ListInterface.CheckIfTitleExists(userID, title); err == nil {
		return errors.New("title don't exist")
	}

	return s.repos.LinkInterface.Update(listID, linkInput)
}

func (s *LinkService) Delete(userID, linkID int, title string) error {
	if err := s.repos.ListInterface.CheckIfTitleExists(userID, title); err == nil {
		return errors.New("title don't exist")
	}

	return s.repos.LinkInterface.Delete(linkID)
}
