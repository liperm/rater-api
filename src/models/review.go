package models

type Review struct {
	Base
	UserID int     `json:"user_id"`
	User   User    `json:"user"`
	Stars  float32 `json:"stars"`
	Text   string  `json:"text"`
	ItemID int     `json:"item_id"`
}

func (r *Review) TableName() string {
	return "rater.review"
}
