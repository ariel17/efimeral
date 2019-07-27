package apierrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ariel17/efimeral/api/config"
	"github.com/getsentry/sentry-go"
)

// APIError is the JSON body to show on a failed response.
type APIError struct {
	Err         error  `json:"-"`
	Description string `json:"description"`
	Cause       string `json:"cause"`
	Status      int    `json:"status"`
}

// NewBadRequestError creates a new APIError for invalid input.
func NewBadRequestError(err error) *APIError {
	apiErr := APIError{
		Err:         err,
		Description: "invalid input data",
		Cause:       err.Error(),
		Status:      http.StatusBadRequest,
	}
	sentry.CaptureException(err)
	return &apiErr
}

// NewNotFoundError creates a new APIError for resource not found.
func NewNotFoundError() *APIError {
	err := errors.New("resource not found")
	apiErr := APIError{
		Err:         err,
		Description: err.Error(),
		Status:      http.StatusNotFound,
	}
	sentry.CaptureException(err)
	return &apiErr
}

// NewInternalServerError creates a new APIError for internal errors.
func NewInternalServerError(err error) *APIError {
	apiErr := APIError{
		Err:         err,
		Description: "internal server error",
		Cause:       err.Error(),
		Status:      http.StatusInternalServerError,
	}
	sentry.CaptureException(err)
	return &apiErr
}

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   config.SentryDSN,
		Debug: config.Environment != config.ProductionEnv,
	})

	if err != nil {
		panic(fmt.Sprintf("Sentry initialization failed: %v", err))
	}
}
