package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/liperm/trabalho_mobile_02/src/encryption"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/handlers"
)

type patchPasswordRequest struct {
	NewPassword string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	id, err := handlers.CreateUser(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.CreateErrorResponse("User", err)
		json.NewEncoder(w).Encode(errorResponse)
		log.Println("[CreateUser] Response ", errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	log.Println("[CreateUser] Response ", response)
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
		log.Println("[GetUserById] Response ", errorResponse)
		return
	}

	user, err := handlers.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := formatters.NotFoundResponse("User")
		json.NewEncoder(w).Encode(errorResponse)
		log.Println("[GetUserById] Response ", errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("[GetUserById] Response ", "OK")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handlers.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := formatters.NotFoundResponse("User")
		json.NewEncoder(w).Encode(errorResponse)
		log.Println("[GetUsers] Response ", errorResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("[GetUsers] Response ", "OK")
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
		log.Println("[GetUpdatePasswordCode] Response ", errorResponse)
		return
	}

	cookie := http.Cookie{
		Name:     "change-password-code",
		Value:    encryption.EncryptData(code),
		Path:     "/users",
		MaxAge:   600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func PatchPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, paramErr := strconv.Atoi(vars["id"])
	if paramErr != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.InvalidParamResponse("id")
		json.NewEncoder(w).Encode(errorResponse)
		log.Println("[PatchPassword] Response ", errorResponse)
		return
	}

	var request patchPasswordRequest
	json.NewDecoder(r.Body).Decode(&request)
	err := handlers.ChangePassword(id, request.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[PatchPassword] Response ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
