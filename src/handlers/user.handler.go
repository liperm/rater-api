package handlers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/liperm/trabalho_mobile_02/src/models"
	"github.com/liperm/trabalho_mobile_02/src/repositories"
)

func CreateUser(requestBody io.ReadCloser) (int, error) {
	var user models.User
	json.NewDecoder(requestBody).Decode(&user)
	err := repositories.CreateUser(&user)

	if err != nil {
		log.Println("Create user fail: ", err)
		return 0, err
	}

	log.Println("New user created: ", user)
	return user.ID, nil
}

func GetUserById(id int) models.User {
	user := repositories.GetUserById(id)

	if user.ID == 0 {
		log.Println("Getting user by ID fail for ID ", id)
		return user
	}

	log.Println("Getting user by ID ", user)
	return user
}
