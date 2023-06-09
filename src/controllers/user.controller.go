package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liperm/trabalho_mobile_02/src/encryption"
	"github.com/liperm/trabalho_mobile_02/src/formatters"
	"github.com/liperm/trabalho_mobile_02/src/handlers"
)

type patchPasswordRequest struct {
	NewPassword string `json:"password"`
}

type ForgotMyPasswordRequest struct {
	Email string `json:"email"`
}

func init() {
	log.SetPrefix("[Controller] ")
}

func CreateUser(c *gin.Context) {
	id, err := handlers.CreateUser(c.Request.Body)

	if err != nil {
		errorResponse := formatters.CreateErrorResponse("User", err)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[CreateUser] Response ", errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	log.Println("[CreateUser] Response ", response)
	c.IndentedJSON(http.StatusCreated, response)
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		errorResponse := formatters.InvalidParamResponse("id")
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[GetUserById] Response ", errorResponse)
		return
	}

	user, err := handlers.GetUserById(id)
	if err != nil {
		errorResponse := formatters.NotFoundResponse("User")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[GetUserById] Response ", errorResponse)
		return
	}
	log.Println("[GetUserById] Response ", "OK")
	c.IndentedJSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	users, err := handlers.GetUsers()
	if err != nil {
		errorResponse := formatters.NotFoundResponse("User")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[GetUsers] Response ", errorResponse)
		return
	}
	log.Println("[GetUsers] Response ", "OK")
	c.IndentedJSON(http.StatusOK, users)
}

func ForgotMyPassword(c *gin.Context) {
	var request ForgotMyPasswordRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorResponse := formatters.InvalidPayloadResponse(err)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[GetUpdatePasswordCode] Response ", errorResponse)
		return
	}

	email := request.Email
	code, err := handlers.SendUpdatePasswordCode(email)
	if err != nil {
		errorResponse := formatters.SendEmailErrorResponse(email, err)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[GetUpdatePasswordCode] Response ", errorResponse)
		return
	}

	cookie := http.Cookie{
		Name:     "change-password-code",
		Value:    encryption.EncryptData(code),
		Path:     "/users",
		MaxAge:   600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, &cookie)
	c.Writer.WriteHeader(http.StatusOK)
}

func PatchPassword(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil || id <= 0 {
		errorResponse := formatters.InvalidParamResponse("id")
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[PatchPassword] Response ", errorResponse)
		return
	}

	var request patchPasswordRequest
	c.BindJSON(&request)
	err := handlers.ChangePassword(id, request.NewPassword)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("[PatchPassword] Response ", err)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func Login(c *gin.Context) {
	id := handlers.Login(c.Request.Body)

	if id == 0 {
		errorResponse := formatters.NotFoundResponse("User")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[Login] Response ", errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	log.Println("[Login] Response ", response)
	c.IndentedJSON(http.StatusCreated, response)
}

func CreateFavorites(c *gin.Context) {
	id, err := handlers.CreateFavorites(c.Request.Body)

	if err != nil {
		errorResponse := formatters.CreateErrorResponse("Favorite", err)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[CreateFavorite] Response ", errorResponse)
		return
	}

	response := formatters.CreateSuccessResponse(id)
	log.Println("[CreateFavorite] Response ", response)
	c.IndentedJSON(http.StatusCreated, response)
}

func DeleteFavorite(c *gin.Context) {
	userId, paramErr := strconv.Atoi(c.Param("userId"))
	itemId, paramErr := strconv.Atoi(c.Param("itemId"))
	if paramErr != nil || userId <= 0 || itemId <= 0 {
		errorResponse := formatters.InvalidParamResponse("id")
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[DeleteFavorite] Response ", errorResponse)
		return
	}

	err := handlers.DeleteFavorite(userId, itemId)
	if err != nil {
		errorResponse := formatters.NotFoundResponse("Favorite")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[DeleteFavorite] Response ", errorResponse)
		return
	}
	log.Println("[DeleteFavorite] Response ", "OK")
	c.Writer.WriteHeader(http.StatusNoContent)
}

func GetFavoritesByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		errorResponse := formatters.InvalidParamResponse("id")
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		log.Println("[GetFavoritesByUserId] Response ", errorResponse)
		return
	}

	favorites, err := handlers.GetFavoritesByUserId(id)
	if err != nil {
		errorResponse := formatters.NotFoundResponse("User")
		c.IndentedJSON(http.StatusNotFound, errorResponse)
		log.Println("[GetFavoritesByUserId] Response ", errorResponse)
		return
	}
	log.Println("[GetFavoritesByUserId] Response ", "OK")
	c.IndentedJSON(http.StatusOK, favorites)
}
