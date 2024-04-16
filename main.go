package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nabinkatwal7/x-backend-go/db"
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
	// db.Database.AutoMigrate(&model.User{})
}
