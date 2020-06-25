package distributions

import "github.com/gin-gonic/gin"

// Routes configures URL mapping for sessions.
func Routes(route *gin.Engine) {
	r := route.Group("/distributions")
	r.GET("/", GetAvailableDistributions)
}
