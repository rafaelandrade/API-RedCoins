package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelandrade/API-RedCoins/api/controllers"
	"github.com/rafaelandrade/API-RedCoins/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error para conseguir arquivo env, %v", err)
	} else {
		fmt.Println("Valores env")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)
	server.Run(":8080")
}