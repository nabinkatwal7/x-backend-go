package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4;primaryKey" json:"postId"`
	UserID   uint      `json:"userId"` // Field to hold the user's ID
	Content  string    `gorm:"type:text" json:"content"`
	Media    string    `gorm:"type:text" json:"media"`
	User     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Likes    []Like    `json:"likes"`
	Comments []Comment `json:"comments"`
}
