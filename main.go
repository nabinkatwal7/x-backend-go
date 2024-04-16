package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}