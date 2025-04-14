package repository

import (
	"cisdi-technical-assessment/REST/data-service/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(page, pageSize int) ([]model.Book, int64, error)
	FindByID(id uint) (*model.Book, error)
	Create(book *model.Book) error
	Update(book *model.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(page, pageSize int) ([]model.Book, int64, error) {
	var books []model.Book
	var total int64

	// Count total records
	if err := r.db.Model(&model.Book{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records with preloaded relations
	offset := (page - 1) * pageSize
	if err := r.db.Preload("Author").Preload("Publisher").Limit(pageSize).Offset(offset).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (r *bookRepository) FindByID(id uint) (*model.Book, error) {
	var book model.Book
	if err := r.db.Preload("Author").Preload("Publisher").First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}
