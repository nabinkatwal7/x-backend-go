package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/x-backend-go/model"
)

func Register(context *gin.Context){
	var input model.RegisterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, savedUser)
}