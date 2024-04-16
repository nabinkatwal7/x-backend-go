package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nabinkatwal7/x-backend-go/db"
	"github.com/nabinkatwal7/x-backend-go/model"
	"github.com/nabinkatwal7/x-backend-go/router"
)

func main() {
	loadEnv()
	loadDatabase()
	router.ServeApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase(){
	db.Connect()
	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.Post{})
	db.Database.AutoMigrate(&model.Comment{})
	db.Database.AutoMigrate(&model.Like{})
	db.Database.AutoMigrate(&model.Follower{})
	db.Database.AutoMigrate(&model.Message{})
	db.Database.AutoMigrate(&model.Notification{})
}
