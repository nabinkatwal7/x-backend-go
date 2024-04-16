package model

import "gorm.io/gorm"

type Comment struct {
    gorm.Model
    UserID  uint   `json:"userId"`
    PostID  uint   `json:"postId"`
    Content string `gorm:"type:text" json:"content"`
    User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
    Post    Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}