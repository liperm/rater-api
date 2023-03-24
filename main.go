package main

import (
	"log"

	"github.com/liperm/trabalho_mobile_02/database"
	"github.com/liperm/trabalho_mobile_02/router"
)

func main() {
	log.Println("Initializing...")
	database.Connect()
	router.HandleRequest()
}
