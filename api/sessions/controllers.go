package sessions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ariel17/efimeral/api/apierrors"
	"github.com/gin-gonic/gin"
)

// Distribution refers to Linux distributions.
type Distribution string

const (
	ubuntu Distribution = "ubuntu"
	arch   Distribution = "arch"
)

var (
	availableDistributions = map[Distribution][]string{
		ubuntu: []string{"19.04"},
		arch:   []string{},
	}
)

// OSDistribution responds to the required data needed to create a new session.
type OSDistribution struct {
	Distribution Distribution `json:"distribution"`
	Tag          string       `json:"tag"`
}

// UnmarshalJSON is a custom implementation for unmarshalling OSDistribution in
// order to check valid distribution name.
func (od *OSDistribution) UnmarshalJSON(data []byte) error {
	type Alias OSDistribution
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if tags, found := availableDistributions[aux.Distribution]; found {
		od.Distribution = aux.Distribution
		var validTag bool
		for _, tag := range tags {
			if aux.Tag == tag {
				validTag = true
			}
		}
		if !validTag {
			return fmt.Errorf("'%s' is not a valid tag for '%s' distribution", aux.Tag, aux.Distribution)
		}
		od.Tag = aux.Tag
	} else {
		return fmt.Errorf("'%s' is not available as distribution", aux.Distribution)
	}
	return nil
}

// CreateSession handles the POST method to create a new session.
func CreateSession(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		apiErr := apierrors.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	var od OSDistribution
	if err := json.Unmarshal(body, &od); err != nil {
		apiErr := apierrors.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	s, apiErr := NewSession(&od)
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
