package model

import "gorm.io/gorm"

type Like struct {
    gorm.Model
    UserID uint `json:"userId"`
    PostID uint `json:"postId"`
    User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
    Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}
