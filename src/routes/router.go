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

	log.Fatal(router.Run("localhost:8080"))
}
