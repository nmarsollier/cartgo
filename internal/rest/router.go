package rest

import (
	"fmt"

	"github.com/nmarsollier/cartgo/internal/engine/env"
	"github.com/nmarsollier/cartgo/internal/rest/engine"
)

// Start this server
func Start() {
	InitRoutes()
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
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
