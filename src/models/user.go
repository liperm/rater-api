package models

type User struct {
	Base
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Advertiser bool   `json:"advertiser"`
}

func (u *User) TableName() string {
	return "rater.user"
}
