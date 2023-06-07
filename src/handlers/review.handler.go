package handlers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/liperm/trabalho_mobile_02/src/models"
	"github.com/liperm/trabalho_mobile_02/src/repositories"
)

func init() {
	log.SetPrefix("[Handler] ")
}

func CreateReview(requestBody io.ReadCloser) (int, error) {
	var review *models.Review
	json.NewDecoder(requestBody).Decode(&review)

	err := repositories.CreateReview(review)
	if err != nil {
		log.Println("Create review fail: ", err)
		return 0, err
	}

	log.Println("New review Created: ", review)
	return review.ID, nil
}
