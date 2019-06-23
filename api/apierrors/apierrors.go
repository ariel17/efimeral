package apierrors

import (
	"fmt"
	"net/http"

	"github.com/ariel17/efimeral/api/config"
	"github.com/getsentry/sentry-go"
)

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
	sentry.CaptureException(err)
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
		Status:      http.StatusInternalServerError,
	}
	sentry.CaptureException(err)
	return &apiErr
}

func init() {
	if config.Environment == config.ProductionEnv {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "https://a07934d966db496eb3a27b003d6e9bfe:49b4d993cc74473bbf9d004edbef7e3c@sentry.io/1488418",
		})

		if err != nil {
			panic(fmt.Sprintf("Sentry initialization failed: %v\n", err))
		}
	}
}
