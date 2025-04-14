package repository

import (
	"cisdi-technical-assessment/REST/data-service/model"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	FindAll(page, pageSize int) ([]model.Publisher, int64, error)
	FindByID(id uint) (*model.Publisher, error)
	Create(publisher *model.Publisher) error
	Update(publisher *model.Publisher) error
	Delete(id uint) error
}

type publisherRepository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) PublisherRepository {
	return &publisherRepository{db}
}

func (r *publisherRepository) FindAll(page, pageSize int) ([]model.Publisher, int64, error) {
	var publishers []model.Publisher
	var total int64

	// Count total records
	if err := r.db.Model(&model.Publisher{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records
	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Find(&publishers).Error; err != nil {
		return nil, 0, err
	}

	return publishers, total, nil
}

func (r *publisherRepository) FindByID(id uint) (*model.Publisher, error) {
	var publisher model.Publisher
	if err := r.db.First(&publisher, id).Error; err != nil {
		return nil, err
	}
	return &publisher, nil
}

func (r *publisherRepository) Create(publisher *model.Publisher) error {
	return r.db.Create(publisher).Error
}

func (r *publisherRepository) Update(publisher *model.Publisher) error {
	return r.db.Save(publisher).Error
}

func (r *publisherRepository) Delete(id uint) error {
	return r.db.Delete(&model.Publisher{}, id).Error
}
