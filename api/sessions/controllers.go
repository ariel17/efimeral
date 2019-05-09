package sessions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSession handles the POST method to create a new session.
func CreateSession(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var sr SessionRequest
	if err := json.Unmarshal(body, &sr); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	s, err := NewSession(&sr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, s)
}

// GetSession handles the GET method to fetch a session data.
func GetSession(c *gin.Context) {
	id := c.Param("id")
	s, err := FetchSession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if s == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, s)
}

// DeleteSession handles the DELETE method to remove an existing session.
func DeleteSession(c *gin.Context) {
	id := c.Param("id")
	s, err := RemoveSession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if s == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, s)
}
