package service

import (
	"cisdi-technical-assessment/REST/data-service/model"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/repository"
	"errors"
	"math"

	"gorm.io/gorm"
)

type BookService interface {
	GetAllBooks(page, pageSize int) (*dto.PaginationResponse, error)
	GetBookByID(id uint) (*model.Book, error)
	CreateBook(req dto.CreateBookRequest) (*model.Book, error)
	UpdateBook(id uint, req dto.UpdateBookRequest) (*model.Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	repo          repository.BookRepository
	authorRepo    repository.AuthorRepository
	publisherRepo repository.PublisherRepository
}

func NewBookService(
	repo repository.BookRepository,
	authorRepo repository.AuthorRepository,
	publisherRepo repository.PublisherRepository,
) BookService {
	return &bookService{repo, authorRepo, publisherRepo}
}

func (s *bookService) GetAllBooks(page, pageSize int) (*dto.PaginationResponse, error) {
	books, total, err := s.repo.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &dto.PaginationResponse{
		TotalRecords: total,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
		Data:         books,
	}, nil
}

func (s *bookService) GetBookByID(id uint) (*model.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(req dto.CreateBookRequest) (*model.Book, error) {
	// Verify that author exists
	_, err := s.authorRepo.FindByID(req.AuthorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}

	// Verify that publisher exists
	_, err = s.publisherRepo.FindByID(req.PublisherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("publisher not found")
		}
		return nil, err
	}

	book := &model.Book{
		Title:           req.Title,
		ISBN:            req.ISBN,
		PublicationYear: req.PublicationYear,
		Description:     req.Description,
		AuthorID:        req.AuthorID,
		PublisherID:     req.PublisherID,
	}

	err = s.repo.Create(book)
	if err != nil {
		return nil, err
	}

	// Reload with relationships
	return s.repo.FindByID(book.ID)
}

func (s *bookService) UpdateBook(id uint, req dto.UpdateBookRequest) (*model.Book, error) {
	// Verify book exists
	book, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	// Verify that author exists
	_, err = s.authorRepo.FindByID(req.AuthorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}

	// Verify that publisher exists
	_, err = s.publisherRepo.FindByID(req.PublisherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("publisher not found")
		}
		return nil, err
	}

	book.Title = req.Title
	book.ISBN = req.ISBN
	book.PublicationYear = req.PublicationYear
	book.Description = req.Description
	book.AuthorID = req.AuthorID
	book.PublisherID = req.PublisherID

	err = s.repo.Update(book)
	if err != nil {
		return nil, err
	}

	// Reload with relationships
	return s.repo.FindByID(book.ID)
}

func (s *bookService) DeleteBook(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
