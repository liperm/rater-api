package repositories

import (
	"github.com/liperm/trabalho_mobile_02/src/database"
	"github.com/liperm/trabalho_mobile_02/src/models"
)

func CreateReview(review *models.Review) error {
	result := database.DB.Create(&review)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
