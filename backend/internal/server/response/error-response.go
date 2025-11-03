package response

import (
	"github.com/Daci1/url-shortener-atad/internal/errs"
	"net/http"
)

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Status      int16  `json:"status"`
	Description string `json:"description"`
}

func NewErrorResponse(status int16, description string) (int, *ErrorResponse) {
	errorResponse := &ErrorResponse{
		Error: ErrorDetail{
			Status:      status,
			Description: description,
		},
	}

	return int(status), errorResponse
}

func NewInternalServerErrorResponse() (int, *ErrorResponse) {
	return NewErrorResponse(http.StatusInternalServerError, "Internal server error.")
}

func NewErrorFromCustomError(err errs.CustomError) (int, *ErrorResponse) {
	return NewErrorResponse(int16(err.Status()), err.Message())
}
