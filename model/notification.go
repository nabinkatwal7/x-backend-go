package model

import "gorm.io/gorm"

type Notification struct {
    gorm.Model
    UserID            uint   `json:"userId"`
    NotificationType  string `json:"notificationType"`
    TargetUserID      uint   `json:"targetUserId,omitempty"`
    TargetPostID      uint   `json:"targetPostId,omitempty"`
    Message           string `gorm:"type:text" json:"message,omitempty"`
    Unseen            bool   `json:"unseen"`
    User              User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
    TargetUser        User   `gorm:"foreignKey:TargetUserID" json:"targetUser,omitempty"`
    TargetPost        Post   `gorm:"foreignKey:TargetPostID" json:"targetPost,omitempty"`
}