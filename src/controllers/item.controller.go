package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/handlers"
	"github.com/liperm/trabalho_mobile_02/src/models"
)

func init() {
	log.SetPrefix("[Controller] ")
}

func CreateItem(c *gin.Context) {
	requestBody := c.Request.Body
	id, err := handlers.CreateItem(requestBody)
	if err != nil {
		errorResponse := formatters.CreateErrorResponse("Item", err)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[CreateItem] Response ", errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	log.Println("[CreateItem] Response ", response)
	c.IndentedJSON(http.StatusCreated, response)
}

func GetItems(c *gin.Context) {
	items, err := handlers.GetItems()
	if err != nil {
		errorResponse := formatters.NotFoundResponse("Items")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[GetItems] Response ", errorResponse)
		return
	}

	log.Println("[GetItems] Response ", "OK")
	c.IndentedJSON(http.StatusOK, items)
}

func GetItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		errorResponse := formatters.InvalidParamResponse("id")
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[GetItemById] Response ", errorResponse)
		return
	}

	item, err := handlers.GetItemById(id)
	if err != nil {
		errorResponse := formatters.NotFoundResponse("Item")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[GetItemById] Response ", errorResponse)
		return
	}
	log.Println("[GetItemById] Response ", "OK")
	c.IndentedJSON(http.StatusOK, item)
}

func GetItemsByCategory(c *gin.Context) {
	category := c.Param("category")
	items, err := handlers.GetItemsByCategory(category)
	if err != nil {
		switch err.Error() {
		case "not found":
			errorResponse := formatters.NotFoundResponse("Item")
			c.IndentedJSON(http.StatusNotFound, errorResponse)
			log.Println("[GetItemsByCategory] Response ", errorResponse)
			return
		case "invalid category":
			errorResponse := formatters.InvalidParamResponse(category)
			c.IndentedJSON(http.StatusNotFound, errorResponse)
			log.Println("[GetItemsByCategory] Response ", errorResponse)
			return
		}
	}
	log.Println("[GetItemsByCategory] Response ", "OK")
	c.IndentedJSON(http.StatusOK, items)
}

func GetCategories(c *gin.Context) {
	var categories []string
	for k := range models.ItemCategory {
		categories = append(categories, k)
	}

	c.IndentedJSON(http.StatusOK, formatters.GetCategoriesResponse(categories))
}

func GetItemsByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		errorResponse := formatters.InvalidParamResponse("id")
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[GetItemByUserId] Response ", errorResponse)
		return
	}

	items, err := handlers.GetItemsByUserId(id)
	if err != nil {
		errorResponse := formatters.NotFoundResponse("Item")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[GetItemByUserId] Response ", errorResponse)
		return
	}
	log.Println("[GetItemByUserId] Response ", "OK")
	c.IndentedJSON(http.StatusOK, items)
}
