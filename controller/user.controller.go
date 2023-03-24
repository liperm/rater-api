package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/liperm/trabalho_mobile_02/formatter"
	"github.com/liperm/trabalho_mobile_02/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := user.Create()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatter.CreateErrorResponse(err)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := formatter.CreateSuccessResponse(user.ID)
	log.Println("New user created", user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
