package sessions

import "github.com/gin-gonic/gin"

// Routes configures URL mapping for sessions.
func Routes(route *gin.Engine) {
	r := route.Group("/sessions")
	r.GET("/:id", GetSession)
	r.POST("", CreateSession)
	r.POST("/", CreateSession)
	r.DELETE("/:id", DeleteSession)
}
