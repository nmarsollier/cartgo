package server

import (
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	_ "github.com/nmarsollier/cartgo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var server *gin.Engine = nil

func Router(deps ...interface{}) *gin.Engine {
	if server == nil {

		server = gin.Default()

		server.Use(gzip.Gzip(gzip.DefaultCompression))
		server.Use(GinLoggerMiddleware(deps...))

		server.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     false,
			ValidateHeaders: false,
		}))

		server.Use(ErrorHandler)

		server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return server
}

// Obtiene el contexto a serivcios externos
// En prod este contexto esta vacio.
func GinDeps(c *gin.Context) []interface{} {
	var deps []interface{}
	// mock_deps solo es para mocks en testing
	if mocks, ok := c.Get("mock_deps"); ok {
		deps = mocks.([]interface{})
	}

	deps = append(deps, ginLogger(c))

	return deps
}
