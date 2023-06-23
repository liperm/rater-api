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
	database.DB.Preload("Reviews").
		Preload("Reviews.User").
		Joins("User").
		Find(&items, "item.active = ?", true)
	return items
}

func GetItemById(id int) models.Item {
	var item models.Item
	database.DB.Preload("Reviews").
		Preload("Reviews.User").
		Joins("User").
		Find(&item, "item.active = ? and item.id = ?", true, id)
	return item
}

func GetItemsByCategory(category string) []models.Item {
	var items []models.Item
	whereClause := "item.active = ? and item.category = ?"
	database.DB.Joins("User").Find(&items, whereClause, true, category)
	return items
}

func UpdateItemAverageRating(i *models.Item, average float32) error {
	result := database.DB.Model(&i).
		Where("active = ?", true).
		Update("average_rating", average)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetItemsByUserId(userId int) []models.Item {
	var item []models.Item
	database.DB.Preload("Reviews").
		Preload("Reviews.User").
		Joins("User").
		Find(&item, "item.active = ? and item.user_id = ?", true, userId)
	return item
}
