package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/handlers"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	id, err := handlers.CreateUser(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.CreateErrorResponse("User", err)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.InvalidParamResponse("id")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user := handlers.GetUserById(id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := formatters.NotFoundResponse("User")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
