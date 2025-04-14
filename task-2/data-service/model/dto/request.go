package dto

import "time"

type CreatePublisherRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" binding:"omitempty,email"`
}

type UpdatePublisherRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" binding:"omitempty,email"`
}

type CreateAuthorRequest struct {
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
}

type UpdateAuthorRequest struct {
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
}

type CreateBookRequest struct {
	Title           string `json:"title" binding:"required"`
	ISBN            string `json:"isbn" binding:"required"`
	PublicationYear int    `json:"publication_year" binding:"required,numeric"`
	Description     string `json:"description"`
	AuthorID        uint   `json:"author_id" binding:"required,numeric"`
	PublisherID     uint   `json:"publisher_id" binding:"required,numeric"`
}

type UpdateBookRequest struct {
	Title           string `json:"title" binding:"required"`
	ISBN            string `json:"isbn" binding:"required"`
	PublicationYear int    `json:"publication_year" binding:"required,numeric"`
	Description     string `json:"description"`
	AuthorID        uint   `json:"author_id" binding:"required,numeric"`
	PublisherID     uint   `json:"publisher_id" binding:"required,numeric"`
}
