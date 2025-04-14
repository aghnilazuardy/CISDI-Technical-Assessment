package service

import (
	"cisdi-technical-assessment/REST/data-service/model"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/repository"
	"errors"
	"math"

	"gorm.io/gorm"
)

type AuthorService interface {
	GetAllAuthors(page, pageSize int) (*dto.PaginationResponse, error)
	GetAuthorByID(id uint) (*model.Author, error)
	CreateAuthor(req dto.CreateAuthorRequest) (*model.Author, error)
	UpdateAuthor(id uint, req dto.UpdateAuthorRequest) (*model.Author, error)
	DeleteAuthor(id uint) error
}

type authorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) GetAllAuthors(page, pageSize int) (*dto.PaginationResponse, error) {
	authors, total, err := s.repo.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &dto.PaginationResponse{
		TotalRecords: total,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
		Data:         authors,
	}, nil
}

func (s *authorService) GetAuthorByID(id uint) (*model.Author, error) {
	return s.repo.FindByID(id)
}

func (s *authorService) CreateAuthor(req dto.CreateAuthorRequest) (*model.Author, error) {
	author := &model.Author{
		Name:      req.Name,
		Biography: req.Biography,
		BirthDate: req.BirthDate,
	}

	err := s.repo.Create(author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (s *authorService) UpdateAuthor(id uint, req dto.UpdateAuthorRequest) (*model.Author, error) {
	author, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}

	author.Name = req.Name
	author.Biography = req.Biography
	author.BirthDate = req.BirthDate

	err = s.repo.Update(author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (s *authorService) DeleteAuthor(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("author not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
