package sessions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/gin-gonic/gin"
)

// CreateSession handles the POST method to create a new session.
func CreateSession(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		apiErr := apierrors.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}

	var sr SessionDistribution
	if err := json.Unmarshal(body, &sr); err != nil {
		apiErr := apierrors.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}

	s, apiErr := NewSession(&sr)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.JSON(http.StatusCreated, s)
}

// GetSession handles the GET method to fetch a session data.
func GetSession(c *gin.Context) {
	id := c.Param("id")
	s, err := FetchSession(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, s)
}

// DeleteSession handles the DELETE method to remove an existing session.
func DeleteSession(c *gin.Context) {
	id := c.Param("id")
	s, err := RemoveSession(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, s)
}
