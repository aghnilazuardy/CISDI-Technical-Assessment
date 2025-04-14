package service

import (
	"cisdi-technical-assessment/REST/data-service/model"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/repository"
	"errors"
	"math"

	"gorm.io/gorm"
)

type PublisherService interface {
	GetAllPublishers(page, pageSize int) (*dto.PaginationResponse, error)
	GetPublisherByID(id uint) (*model.Publisher, error)
	CreatePublisher(req dto.CreatePublisherRequest) (*model.Publisher, error)
	UpdatePublisher(id uint, req dto.UpdatePublisherRequest) (*model.Publisher, error)
	DeletePublisher(id uint) error
}

type publisherService struct {
	repo repository.PublisherRepository
}

func NewPublisherService(repo repository.PublisherRepository) PublisherService {
	return &publisherService{repo}
}

func (s *publisherService) GetAllPublishers(page, pageSize int) (*dto.PaginationResponse, error) {
	publishers, total, err := s.repo.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &dto.PaginationResponse{
		TotalRecords: total,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
		Data:         publishers,
	}, nil
}

func (s *publisherService) GetPublisherByID(id uint) (*model.Publisher, error) {
	return s.repo.FindByID(id)
}

func (s *publisherService) CreatePublisher(req dto.CreatePublisherRequest) (*model.Publisher, error) {
	publisher := &model.Publisher{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}

	err := s.repo.Create(publisher)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func (s *publisherService) UpdatePublisher(id uint, req dto.UpdatePublisherRequest) (*model.Publisher, error) {
	publisher, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("publisher not found")
		}
		return nil, err
	}

	publisher.Name = req.Name
	publisher.Address = req.Address
	publisher.Phone = req.Phone
	publisher.Email = req.Email

	err = s.repo.Update(publisher)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func (s *publisherService) DeletePublisher(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("publisher not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
