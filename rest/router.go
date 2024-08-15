package rest

import (
	"fmt"

	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/tools/env"
)

// Start this server
func Start() {
	InitRoutes()
	server.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}

func InitRoutes() {
	initDeleteCart()
	initGetCartValidate()
	initGetCart()
	initPostCartArticleDecrement()
	initPostCartArticleIncrement()
	initPostCartArticle()
	initPostCartCheckout()
}
