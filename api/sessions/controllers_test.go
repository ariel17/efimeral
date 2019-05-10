package sessions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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

func TestGetSession(t *testing.T) {
	testCases := []struct {
		name   string
		status int
		body   string
	}{
		{"existing session", http.StatusOK, `{}`},
		{"fails to fetch session", http.StatusInternalServerError, `{}`},
		{"not found", http.StatusNotFound, `{}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/sessions/fake-id", nil)
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
		{"creates new session", `{}`, http.StatusCreated, `{}`},
		{"invalid distribution", `{}`, http.StatusBadRequest, `{}`},
		{"fails to create new session", `{}`, http.StatusInternalServerError, `{}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer([]byte(tc.data)))
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.body, w.Body.String())
		})
	}
}

func TestDeleteSession(t *testing.T) {
	testCases := []struct {
		name   string
		data   string
		status int
		body   string
	}{
		{"deletes existing session", `{}`, http.StatusOK, `{}`},
		{"session not found", `{}`, http.StatusNotFound, `{}`},
		{"fails to delete session", `{}`, http.StatusInternalServerError, `{}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/sessions/fake-id", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.body, w.Body.String())
		})
	}
}
