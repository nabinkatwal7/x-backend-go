package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nabinkatwal7/x-backend-go/db"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase(){
	db.Connect()
	// db.Database.AutoMigrate(&model.User{})
}