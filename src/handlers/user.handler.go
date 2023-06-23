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

	"github.com/liperm/trabalho_mobile_02/src/encryption"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/interfaces"
	"github.com/liperm/trabalho_mobile_02/src/models"
	"github.com/liperm/trabalho_mobile_02/src/repositories"
	"gopkg.in/gomail.v2"
)

var (
	targetEmail   string = os.Getenv("TARGET_EMAIL")
	email         string = os.Getenv("EMAIL")
	emailPassword string = os.Getenv("EMAIL_PASSWORD")
)

func init() {
	log.SetPrefix("[Handler] ")
}

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

func GetUserById(id int) (models.User, error) {
	user := repositories.GetUserById(id)

	if user.ID == 0 {
		log.Println("Get user by ID fail for ID ", id)
		return user, errors.New("not found")
	}

	log.Println("Get User By ID: ", user)
	return user, nil
}

func GetUsers() ([]models.User, error) {
	users := repositories.GetUsers()

	if len(users) == 0 {
		log.Println("Get users fail")
		return users, errors.New("not found")
	}

	log.Println("Get Users ", users)
	return users, nil
}

func SendUpdatePasswordCode(email string) (string, error) {
	user := repositories.GetUserByEmail(email)
	if user.ID <= 0 {
		log.Println("Get user by Email fail for Email ", email, " not found")
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
		log.Println("[sendMail] Error: ", err)
		return err
	}

	log.Println("[sendmail] Email sent to ", userEmail)
	return nil
}

func ChangePassword(id int, newPassword string) error {
	if newPassword == "" {
		return errors.New("invalid password")
	}
	user := repositories.GetUserById(id)
	if user.ID == 0 {
		log.Println("User not found for ID ", id)
		return errors.New("not found")
	}

	err := repositories.UpdatePassword(&user, newPassword)
	if err != nil {
		log.Println("Update Password error: ", err)
		return err
	}

	return nil
}

func Login(requestBody io.ReadCloser) int {
	var request interfaces.Login
	json.NewDecoder(requestBody).Decode(&request)
	user := repositories.GetUserByEmail(request.Email)

	if user.ID == 0 {
		log.Println("User not found")
		return 0
	}

	hashedPassword := encryption.EncryptData(request.Password)
	if user.Password == hashedPassword {
		return user.ID
	}

	return 0
}

func CreateFavorites(requestBody io.ReadCloser) (int, error) {
	var favorite models.Favorites
	err := json.NewDecoder(requestBody).Decode(&favorite)
	if err != nil {
		log.Println("Request Error")
		return 0, errors.New("invalid_request")
	}

	id, err := repositories.CreateFavorites(&favorite)
	if err != nil {
		log.Println("Error creating favorites")
		return 0, err
	}

	return id, nil
}

func DeleteFavorite(userId int, itemId int) error {
	err := repositories.DeleteFavorites(userId, itemId)
	if err != nil {
		log.Println("Error deleting favorites")
		return err
	}

	return nil
}

func GetFavoritesByUserId(id int) ([]models.Item, error) {
	favorites, err := repositories.GetFavoritesByUserId(id)
	var items []models.Item
	if err != nil {
		log.Println("Get Favorites By User ID fail", id)
		return items, errors.New("not found")
	}

	items = formatters.GetFavoriteItemsResponse(favorites)

	log.Println("Get Favorites By User ID: ", favorites)
	return items, nil
}
