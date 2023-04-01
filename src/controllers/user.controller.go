package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/models"
	"github.com/liperm/trabalho_mobile_02/src/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := repositories.CreateUser(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.CreateErrorResponse(err)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(user.ID)
	log.Println("New user created", user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
