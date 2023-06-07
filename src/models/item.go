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

var ItemCategory = map[string]itemCategory{
	"electronic": Eletronic,
	"book":       Book,
	"furniture":  Furniture,
	"video_game": VideoGame,
	"board_game": BoardGame,
	"clothe":     Clothe,
	"vehicle":    Vehicle,
}

type Item struct {
	Base
	Name          string       `json:"name"`
	Category      itemCategory `json:"category"`
	AverageRating float32      `json:"average_rating"`
	Price         float32      `json:"price"`
	BrandName     string       `json:"brand_name"`
	UserID        int          `json:"user_id"`
	User          User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
	Reviews       []Review     `json:"reviews"`
}

func (i *Item) TableName() string {
	return "rater.item"
}
