package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/handlers"
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
