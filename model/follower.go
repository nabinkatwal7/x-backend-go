package model

import "gorm.io/gorm"

type Follower struct {
    gorm.Model
    FollowerID  uint `json:"followerId"`
    FollowingID uint `json:"followingId"`
    Follower    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"follower"`
    Following   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"following"`
}
