package formatter

type createErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type createSuccessResponse struct {
	ID int `json:"id"`
}

func CreateErrorResponse(err error) createErrorResponse {
	return createErrorResponse{
		Message: "an error occurred while creating a new entry",
		Error:   err.Error(),
	}
}

func CreateSuccessResponse(id int) createSuccessResponse {
	return createSuccessResponse{ID: id}
}
