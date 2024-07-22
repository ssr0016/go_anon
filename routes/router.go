package routes

import (
	"gotest/config"

	"github.com/gin-gonic/gin"
)

type Router interface {
	gin.IRouter
	Serve() error
}

type router struct {
	*gin.Engine
	c *config.Config
}

func NewRouter(c *config.Config) Router {
	config := c.Get()
	r := gin.New()

	if config.GetString("ENVIRONMENT") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.GetBool("app.log") {
		r.Use(gin.Logger())
	}

	return &router{Engine: r, c: c}
}

func (r *router) Serve() error {
	port := r.c.Get().GetString("app.port")
	return r.Run(":" + port)
}
