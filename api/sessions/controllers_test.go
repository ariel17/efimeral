package sessions

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/ariel17/efimeral/api/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	r = gin.Default()
)

func TestMain(m *testing.M) {
	Routes(r)
	code := m.Run()
	os.Exit(code)
}

func TestSessionDistribution(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mustFail bool
	}{
		{"valid for ubuntu", `{"distribution":"ubuntu","tag":"19.04"}`, false},
		// {"valid for arch", `{"distribution":"arch"}`, false},
		// {"invalid for mint", `{"distribution":"mint"}`, true},
		{"invalid JSON", `{"distribution":"`, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var d OSDistribution
			err := json.Unmarshal([]byte(tc.input), &d)
			if tc.mustFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetSession(t *testing.T) {
	id := "fake-id"

	testCases := []struct {
		name   string
		status int
		body   string
	}{
		{"existing session", http.StatusOK, `{"id":"fake-id","distribution":"ubuntu","tag":"19.04","url":"http://localhost:9000","created_at":"2019-01-01T00:00:00Z"}`},
		{"fails to fetch session", http.StatusInternalServerError, `{"description":"internal server error","cause":"mocked error","status":500}`},
		{"not found", http.StatusNotFound, `{"description":"resource not found","cause":"","status":404}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dm := newDockerMock()
			if tc.status == http.StatusOK {
				dm.c = &Container{
					ID:           id,
					Distribution: ubuntu,
					Tag:          "19.04",
					CreatedAt:    config.Now().UTC(),
					URL:          "http://localhost:9000",
				}

			} else if tc.status == http.StatusInternalServerError {
				dm.err = apierrors.NewInternalServerError(errors.New("mocked error"))
			} else if tc.status == http.StatusNotFound {
				dm.err = apierrors.NewNotFoundError()
			}
			dc = dm

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/sessions/"+id, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.body, w.Body.String())
		})
	}
}

func TestCreateSession(t *testing.T) {
	testCases := []struct {
		name   string
		data   string
		status int
		body   string
	}{
		{"creates new session", `{"distribution":"ubuntu","tag":"19.04"}`, http.StatusCreated, `{"id":"random-id","distribution":"ubuntu","tag":"19.04","url":"http://localhost:9000","created_at":"2019-01-01T00:00:00Z"}`},
		{"invalid distribution", `{"distribution":"xxx","tag":"yyy"}`, http.StatusBadRequest, `{"description":"invalid input data","cause":"'xxx' is not available as distribution","status":400}`},
		{"invalid tag", `{"distribution":"ubuntu","tag":"yyy"}`, http.StatusBadRequest, `{"description":"invalid input data","cause":"'yyy' is not a valid tag for 'ubuntu' distribution","status":400}`},
		{"fails to create new session", `{"distribution":"ubuntu","tag":"19.04"}`, http.StatusInternalServerError, `{"description":"internal server error","cause":"mocked error","status":500}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dm := newDockerMock()
			if tc.status == http.StatusInternalServerError {
				dm.err = apierrors.NewInternalServerError(errors.New("mocked error"))
			}
			dc = dm

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer([]byte(tc.data)))
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.body, w.Body.String())
		})
	}
}

func TestDeleteSession(t *testing.T) {
	id := "fake-id"

	testCases := []struct {
		name   string
		data   string
		status int
		body   string
	}{
		{"deletes existing session", `{}`, http.StatusOK, `{"id":"fake-id","distribution":"ubuntu","tag":"19.04","url":"http://localhost:9000","created_at":"2019-01-01T00:00:00Z","deleted_at":"2019-01-01T00:00:00Z"}`},
		{"session not found", `{}`, http.StatusNotFound, `{"description":"resource not found","cause":"","status":404}`},
		{"fails to delete session", `{}`, http.StatusInternalServerError, `{"description":"internal server error","cause":"mocked error","status":500}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dm := newDockerMock()
			if tc.status == http.StatusOK {
				dm.c = &Container{
					ID:           id,
					Distribution: ubuntu,
					Tag:          "19.04",
					CreatedAt:    config.Now().UTC(),
					URL:          "http://localhost:9000",
				}

			} else if tc.status == http.StatusInternalServerError {
				dm.err = apierrors.NewInternalServerError(errors.New("mocked error"))
			} else if tc.status == http.StatusNotFound {
				dm.err = apierrors.NewNotFoundError()
			}
			dc = dm

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/sessions/fake-id", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.body, w.Body.String())
		})
	}
}
