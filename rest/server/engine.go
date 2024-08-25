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

func Router(ctx ...interface{}) *gin.Engine {
	if server == nil {

		server = gin.Default()

		server.Use(gzip.Gzip(gzip.DefaultCompression))
		server.Use(GinLoggerMiddleware(ctx...))

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
func GinCtx(c *gin.Context) []interface{} {
	var ctx []interface{}
	// mock_ctx solo es para mocks en testing
	if mocks, ok := c.Get("mock_ctx"); ok {
		ctx = mocks.([]interface{})
	}

	ctx = append(ctx, ginLogger(c))

	return ctx
}
