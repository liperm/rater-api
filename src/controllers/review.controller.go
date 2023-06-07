package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/handlers"
)

func CreateReview(c *gin.Context) {
	requestBody := c.Request.Body
	id, err := handlers.CreateReview(requestBody)
	if err != nil {
		errorResponse := formatters.CreateErrorResponse("Review", err)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[CreateReview] Response ", errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	log.Println("[CreateReview] Response ", response)
	c.IndentedJSON(http.StatusCreated, response)
}
