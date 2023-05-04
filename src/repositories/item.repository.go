package repositories

import (
	"github.com/liperm/trabalho_mobile_02/src/database"
	"github.com/liperm/trabalho_mobile_02/src/models"
)

func CreateItem(item *models.Item) error {
	result := database.DB.Create(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetItems() []models.Item {
	var items []models.Item
	database.DB.Joins("User").Find(&items, "item.active = ?", true)
	return items
}

func GetItemById(id int) models.Item {
	var item models.Item
	database.DB.Joins("User").Find(&item, "item.active = ? and item.id = ?", true, id)
	return item
}

func GetItemsByCategory(category string) []models.Item {
	var items []models.Item
	whereClause := "item.active = ? and item.category = ?"
	database.DB.Joins("User").Find(&items, whereClause, true, category)
	return items
}
