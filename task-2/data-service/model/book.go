package model

import (
	"time"
)

type Book struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"size:200;not null" json:"title"`
	ISBN            string    `gorm:"size:20;unique" json:"isbn"`
	PublicationYear int       `json:"publication_year"`
	Description     string    `gorm:"type:text" json:"description"`
	AuthorID        uint      `json:"author_id"`
	PublisherID     uint      `json:"publisher_id"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Author          Author    `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Publisher       Publisher `gorm:"foreignKey:PublisherID" json:"publisher,omitempty"`
}
