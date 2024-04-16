package model

import "gorm.io/gorm"

type Message struct {
    gorm.Model
    SenderID   uint   `json:"senderId"`
    RecipientID uint   `json:"recipientId"`
    Message    string `gorm:"type:text" json:"message"`
    Seen       bool   `json:"seen"`
    Sender     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sender"`
    Recipient  User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"recipient"`
}