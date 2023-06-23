package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liperm/trabalho_mobile_02/src/controllers"
)

func HandleRequest() {
	router := gin.Default()

	// USERS
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUserById)
	router.GET("/users", controllers.GetUsers)
	router.PATCH("/auth/change-password/:token", controllers.PatchPassword)
	router.POST("/auth/change-password", controllers.ForgotMyPassword)
	router.POST("/auth/login", controllers.Login)
	router.POST("/users/favorites", controllers.CreateFavorites)
	router.DELETE("/users/:userId/favorites/:itemId", controllers.DeleteFavorite)
	router.GET("/users/:id/favorites", controllers.GetFavoritesByUserId)
	router.GET("/users/:id/items", controllers.GetItemsByUserId)

	// ITEMS
	router.POST("/items", controllers.CreateItem)
	router.GET("/items", controllers.GetItems)
	router.GET("/items/:id", controllers.GetItemById)
	router.GET("/items/category/:category", controllers.GetItemsByCategory)
	router.POST("/reviews", controllers.CreateReview)

	// CATEGORY
	router.GET("/categories", controllers.GetCategories)

	log.Fatal(router.Run("localhost:8080"))
}
