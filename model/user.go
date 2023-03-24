package model

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/liperm/trabalho_mobile_02/database"
)

type User struct {
	Base
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Advertiser bool   `json:"advertiser"`
}

func (u *User) TableName() string {
	return "rater.customer"
}

func (u *User) encriptPassword() {
	passord := u.Password
	encryption := sha256.New()
	encryption.Write([]byte(passord))
	u.Password = hex.EncodeToString(encryption.Sum(nil))
}

func (u *User) Create() error {
	u.encriptPassword()
	result := database.DB.Create(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
