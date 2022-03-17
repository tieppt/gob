package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(255);uniqueIndex"`
	Email    string    `gorm:"type:varchar(255);uniqueIndex"`
	Password string    `gorm:"type:text" json:"-"`
	Posts    []Post
}
