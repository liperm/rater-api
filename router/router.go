package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liperm/trabalho_mobile_02/controller"
)

func HandleRequest() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", controller.CreateUser).Methods("POST")

	log.Println("Listenning and serving at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
