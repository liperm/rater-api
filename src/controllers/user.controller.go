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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := handlers.GetUsers()
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := formatters.NotFoundResponse("User")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUpdatePasswordCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	code, err := handlers.SendUpdatePasswordCode(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.SendEmailErrorResponse(email, err)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	cookie := http.Cookie{
		Name:     "change-password-code",
		Value:    code,
		Path:     "/users",
		MaxAge:   600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
