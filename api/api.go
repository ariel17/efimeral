package main

import (
	"github.com/ariel17/efimeral/api/config"
	"github.com/ariel17/efimeral/api/distributions"
	"github.com/ariel17/efimeral/api/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	distributions.Routes(r)
	sessions.Routes(r)
	sessions.RunCleaner()

	if config.Environment == config.ProductionEnv {
		if err := sessions.PullAvailableImages(); err != nil {
			panic(err.Err)
		}
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
