package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/liperm/trabalho_mobile_02/src/database"
	"github.com/liperm/trabalho_mobile_02/src/routes"
)

func main() {
	log.Println("Initializing...")
	database.Connect()
	routes.HandleRequest()
}
