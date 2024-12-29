package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/env"
	"github.com/nmarsollier/cartgo/internal/rest/server"
)

// Start this server
func Start() {
	engine := server.Router()
	InitRoutes(engine)
	engine.Run(fmt.Sprintf(":%d", env.Get().Port))
}

func InitRoutes(engine *gin.Engine) {
	initDeleteCart(engine)
	initGetCartValidate(engine)
	initGetCart(engine)
	initPostCartArticleDecrement(engine)
	initPostCartArticleIncrement(engine)
	initPostCartArticle(engine)
	initPostCartCheckout(engine)
}
