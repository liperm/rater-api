package models

type Favorites struct {
	Base
	User   User `json:"user,omitempty"`
	Item   Item `json:"item"`
	UserID int  `json:"user_id" gorm:""`
	ItemID int  `json:"item_id"`
}

func (f *Favorites) TableName() string {
	return "rater.favorite"
}
