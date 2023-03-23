package main

import (
	"log"

	"github.com/liperm/trabalho_mobile_02/database"
)

func main() {
	log.Println("Initializing...")
	database.Connect()
}
