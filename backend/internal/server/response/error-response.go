package response

import "net/http"

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Status      int16  `json:"status"`
	Description string `json:"description"`
}

func NewErrorResponse(status int16, description string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorDetail{
			Status:      status,
			Description: description,
		},
	}
}

func NewInternalServerErrorResponse() *ErrorResponse {
	return NewErrorResponse(http.StatusInternalServerError, "Internal server error.")
}
