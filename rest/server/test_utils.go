package server

import (
	"github.com/gin-gonic/gin"
)

// Obtiene Router server con el contexto de testing adecuado
// mockeando interfaces a serivcios externos
func TestRouter(ctx ...interface{}) *gin.Engine {
	server = nil
	Router()
	if len(ctx) > 0 {
		server.Use(func(c *gin.Context) {
			c.Set("mock_ctx", ctx)
			c.Next()
		})
	}
	return server
}

// Obtiene el contexto de interfaces mockeadas a serivcios externos
// En prod este contexto esta vacio.
func TestCtx(c *gin.Context) []interface{} {
	var ctx []interface{}
	if mocks, ok := c.Get("mock_ctx"); ok {
		ctx = mocks.([]interface{})
	}
	return ctx
}
