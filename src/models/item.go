package models

type itemCategory string

const (
	Eletronic itemCategory = "eletronic"
	Book      itemCategory = "book"
	Furniture itemCategory = "furniture"
	VideoGame itemCategory = "video_game"
	BoardGame itemCategory = "board_game"
	Clothe    itemCategory = "clothe"
	Vehicle   itemCategory = "vehicle"
)

type Item struct {
	Base
	Name          string       `json:"name"`
	Category      itemCategory `json:"category"`
	AverageRating float32      `json:"average_rating"`
	Price         float32      `json:"price"`
	BrandName     string       `json:"brand_name"`
	UserID        int          `json:"user_id"`
	User          User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
}

func (i *Item) TableName() string {
	return "rater.item"
}
