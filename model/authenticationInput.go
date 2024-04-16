package model

import "github.com/google/uuid"

type RegisterInput struct {
	UserID   uuid.UUID `json:"userId"`
	Email    string    `json:"email" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
