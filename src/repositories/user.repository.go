package repositories

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/liperm/trabalho_mobile_02/src/database"
	"github.com/liperm/trabalho_mobile_02/src/models"
)

func encriptPassword(u *models.User) {
	passord := u.Password
	encryption := sha256.New()
	encryption.Write([]byte(passord))
	u.Password = hex.EncodeToString(encryption.Sum(nil))
}

func CreateUser(u *models.User) error {
	encriptPassword(u)
	result := database.DB.Create(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserById(id int) models.User {
	var user models.User
	database.DB.Where("active = ?", true).Find(&user, id)
	return user
}

func GetUsers() []models.User {
	var users []models.User
	database.DB.Where("active = ?", true).Find(&users)
	return users
}

func GetUserByEmail(email string) models.User {
	var user models.User
	database.DB.Where("active = ? AND email = ?", true, email).Find(&user)
	return user
}
