package apierrors

import "net/http"

// APIError is the JSON body to show on a failed response.
type APIError struct {
	Description string `json:"description"`
	Cause       string `json:"cause"`
	Status      int    `json:"status"`
}

// NewBadRequestError creates a new APIError for invalid input.
func NewBadRequestError(err error) *APIError {
	apiErr := APIError{
		Description: "invalid input data",
		Cause:       err.Error(),
		Status:      http.StatusBadRequest,
	}
	return &apiErr
}

// NewNotFoundError creates a new APIError for resource not found.
func NewNotFoundError() *APIError {
	apiErr := APIError{
		Description: "resource not found",
		Status:      http.StatusNotFound,
	}
	return &apiErr
}

// NewInternalServerError creates a new APIError for internal errors.
func NewInternalServerError(err error) *APIError {
	apiErr := APIError{
		Description: "internal server error",
		Cause:       err.Error(),
		Status:      http.StatusNotFound,
	}
	return &apiErr
}
