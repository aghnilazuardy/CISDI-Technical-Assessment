package model

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Biography string    `gorm:"type:text" json:"biography"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Books     []Book    `gorm:"foreignKey:AuthorID" json:"books,omitempty"`
}
