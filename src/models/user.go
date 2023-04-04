package models

type User struct {
	Base
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture" gorm:"default: 'https://drive.google.com/file/d/1QX9Zhj1eFTjDOSqQIRpBCm8IHjwPO8om/view?usp=share_link'"`
}

func (u *User) TableName() string {
	return "rater.user"
}
