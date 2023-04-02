package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liperm/trabalho_mobile_02/src/controllers"
)

func HandleRequest() {
	router := mux.NewRouter()

	// USERS
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{email}/change-password/code", controllers.GetUpdatePasswordCode).Methods("GET")
	router.HandleFunc("/users/{id}/change-password", controllers.PatchPassword).Methods("PATCH")

	log.Println("Listenning and serving at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
