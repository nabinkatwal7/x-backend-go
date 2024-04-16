package model

import (
	"html"
	"strings"

	// "github.com/google/uuid"
	"github.com/nabinkatwal7/x-backend-go/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID           string         `gorm:"type:uuid;default:uuid_generate_v4;primaryKey" json:"userId"`
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

func (user *User) Save() (*User, error) {
	err := db.Database.Create(&user).Error

	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}
