package repository

import (
	"cisdi-technical-assessment/REST/data-service/model"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll(page, pageSize int) ([]model.Author, int64, error)
	FindByID(id uint) (*model.Author, error)
	Create(author *model.Author) error
	Update(author *model.Author) error
	Delete(id uint) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) FindAll(page, pageSize int) ([]model.Author, int64, error) {
	var authors []model.Author
	var total int64

	// Count total records
	if err := r.db.Model(&model.Author{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records
	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Find(&authors).Error; err != nil {
		return nil, 0, err
	}

	return authors, total, nil
}

func (r *authorRepository) FindByID(id uint) (*model.Author, error) {
	var author model.Author
	if err := r.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *authorRepository) Create(author *model.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepository) Update(author *model.Author) error {
	return r.db.Save(author).Error
}

func (r *authorRepository) Delete(id uint) error {
	return r.db.Delete(&model.Author{}, id).Error
}
