package distributions

import (
	"net/http"

	"github.com/ariel17/efimeral/api/config"
	"github.com/gin-gonic/gin"
)

// GetAvailableDistributions Shows the list with tags for available
// distributions.
func GetAvailableDistributions(c *gin.Context) {
	c.JSON(http.StatusOK, config.AvailableDistributions)
}
