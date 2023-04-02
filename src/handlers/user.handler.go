package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/liperm/trabalho_mobile_02/src/models"
	"github.com/liperm/trabalho_mobile_02/src/repositories"
	"gopkg.in/gomail.v2"
)

var (
	targetEmail   string = os.Getenv("TARGET_EMAIL")
	email         string = os.Getenv("EMAIL")
	emailPassword string = os.Getenv("EMAIL_PASSWORD")
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
		log.Println("Get user by ID fail for ID ", id)
		return user
	}

	log.Println("Get User By ID: ", user)
	return user
}

func GetUsers() []models.User {
	users := repositories.GetUsers()

	if len(users) == 0 {
		log.Println("Get users fail")
		return users
	}

	log.Println("Get Users ", users)
	return users
}

func SendUpdatePasswordCode(email string) (string, error) {
	user := repositories.GetUserByEmail(email)
	if user.ID <= 0 {
		log.Println("Get user by Email fail for Email ", email)
		return "", errors.New("user not found")
	}

	code := generateRandomString(6)
	err := sendMail(email, code)
	if err != nil {
		return "", err
	}
	return code, nil
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func sendMail(userEmail string, code string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", email)
	if targetEmail != "" {
		m.SetHeader("To", targetEmail)
	} else {
		m.SetHeader("To", userEmail)
	}
	m.SetHeader("Subject", "Change Password Request")
	m.SetBody("text/plain", fmt.Sprint("Hi!\nHere is your authorization code to change your password: "+code))

	d := gomail.NewDialer("smtp.gmail.com", 587, email, emailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
