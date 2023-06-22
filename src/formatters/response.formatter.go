package formatters

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error, omitempty"`
}

type createSuccessResponse struct {
	ID int `json:"id"`
}

func CreateErrorResponse(modelName string, err error) errorResponse {
	return errorResponse{
		Message: "an error occurred while creating a new entry for " + modelName,
		Error:   err.Error(),
	}
}

func CreateSuccessResponse(id int) createSuccessResponse {
	return createSuccessResponse{ID: id}
}

func InvalidParamResponse(param string) errorResponse {
	return errorResponse{
		Message: "Invalid parameter found in request. Param: " + param,
	}
}

func NotFoundResponse(modelName string) errorResponse {
	return errorResponse{
		Message: modelName + " not found",
	}
}

func SendEmailErrorResponse(email string, err error) errorResponse {
	return errorResponse{
		Message: "an error occurred while sending email to " + email,
		Error:   err.Error(),
	}
}

func CookieNotFoundResponse(err error) errorResponse {
	return errorResponse{
		Message: "could not find cookie",
		Error:   err.Error(),
	}
}

func InvalidPayloadResponse(err error) errorResponse {
	return errorResponse{
		Message: "invalid payload",
		Error:   err.Error(),
	}
}

type category struct {
	Name string `json:"name"`
}
type getCategoriesResponse struct {
	Categories []category `json:"categories"`
}

func GetCategoriesResponse(categoryNames []string) getCategoriesResponse {
	var formattedCategories []category
	for _, c := range categoryNames {
		category := category{Name: c}
		formattedCategories = append(formattedCategories, category)
	}
	return getCategoriesResponse{Categories: formattedCategories}
}
