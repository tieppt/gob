package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID uuid.UUID `gorm:"type:uuid"`
	Title  string    `gorm:"type:text"`
	Body   string    `gorm:"type:text"` // this is parsed HTML content
	Raw    string    `gorm:"type:text"` // this is raw Markdown content
	User   User
}
