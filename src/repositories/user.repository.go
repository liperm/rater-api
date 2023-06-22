package repositories

import (
	"github.com/liperm/trabalho_mobile_02/src/database"
	"github.com/liperm/trabalho_mobile_02/src/encryption"
	"github.com/liperm/trabalho_mobile_02/src/models"
)

func CreateUser(u *models.User) error {
	u.Password = encryption.EncryptData(u.Password)
	result := database.DB.Create(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserById(id int) models.User {
	var user models.User
	database.DB.Preload("Favorites", "active = ?", true).
		Preload("Favorites.Item").
		Find(&user, "active = ? and id = ?", true, id)
	return user
}

func GetUsers() []models.User {
	var users []models.User
	database.DB.Preload("Favorites").
		Preload("Favorites.Item").
		Find(&users, "active = ?", true)
	return users
}

func GetUserByEmail(email string) models.User {
	var user models.User
	database.DB.
		Where("active = ? AND email = ?", true, email).
		Find(&user)
	return user
}

func UpdatePassword(u *models.User, newPassword string) error {
	encryptedPassword := encryption.EncryptData(newPassword)
	result := database.DB.Model(&u).
		Where("active = ?", true).
		Update("password", encryptedPassword)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CreateFavorites(f *models.Favorites) (int, error) {
	result := database.DB.Create(&f)
	if result.Error != nil {
		return 0, result.Error
	}

	return f.ID, nil
}

func GetFavoriteByID(id int) models.Favorites {
	var f models.Favorites
	database.DB.First(&f, id)

	return f
}

func DeleteFavorites(f *models.Favorites) error {
	result := database.DB.Model(&f).
		Update("active = ?", false)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
