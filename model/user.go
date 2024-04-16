package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4;primaryKey" json:"userId"`
	Name             string         `gorm:"size:255;not null;unique;" json:"name"`
	Username         string         `gorm:"size:255;not null;unique;" json:"username"`
	Email            string         `gorm:"size:255;not null;unique;" json:"email"`
	Password         string         `gorm:"size:255;not null;" json:"password"`
	Bio              string         `gorm:"size:255;" json:"bio"`
	ProfilePicture   []byte         `gorm:"not null;" json:"profile_picture"`
	Posts            []Post         `gorm:"foreignKey:UserID"`
	Followers        []Follower     `gorm:"foreignKey:FollowerID"`  // Followers relationship
	Following        []Follower     `gorm:"foreignKey:FollowingID"` // Following relationship
	Likes            []Like         `json:"likes"`
	Comments         []Comment      `json:"comments"`
	SentMessages     []Message      `gorm:"foreignKey:SenderID"`    // Relationship with sent messages
	ReceivedMessages []Message      `gorm:"foreignKey:RecipientID"` // Relationship with received messages
	Notifications    []Notification `json:"notifications"`
}
