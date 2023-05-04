package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/liperm/trabalho_mobile_02/src/models"
	"github.com/liperm/trabalho_mobile_02/src/repositories"
)

func init() {
	log.SetPrefix("[Handler] ")
}

func CreateItem(requestBody io.ReadCloser) (int, error) {
	var item models.Item
	json.NewDecoder(requestBody).Decode(&item)

	err := repositories.CreateItem(&item)
	if err != nil {
		log.Println("Create item fail: ", err)
		return 0, err
	}

	log.Println("New Item Created: ", item)
	return item.ID, nil
}

func GetItems() ([]models.Item, error) {
	items := repositories.GetItems()

	if len(items) == 0 {
		log.Println("Get items fail")
		return items, errors.New("not found")
	}

	log.Println("Get Items ", items)
	return items, nil
}

func GetItemById(id int) (models.Item, error) {
	item := repositories.GetItemById(id)
	if item.ID <= 0 {
		log.Println("Get item by ID fail for ID ", id)
		return item, errors.New("not found")
	}

	log.Println("Get Item By ID: ", item)
	return item, nil
}

func GetItemsByCategory(category string) ([]models.Item, error) {
	if !isCategoryValid(category) {
		return nil, errors.New("invalid category")
	}

	items := repositories.GetItemsByCategory(category)
	if len(items) == 0 {
		log.Println("Get items fail")
		return nil, errors.New("not found")

	}

	log.Println("Get Items By Category ", items)
	return items, nil
}

func isCategoryValid(category string) bool {
	_, isValid := models.ItemCategory[category]
	return isValid
}
