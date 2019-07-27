package main

import (
	"github.com/ariel17/efimeral/api/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	sessions.Routes(r)
	sessions.RunCleaner()
	r.Run()
}
