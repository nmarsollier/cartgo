package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/rest/middlewares"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// Agrega un articulo del carrito actual.
// @Summary		Agregar Articulo
// @Description	Agrega un articulo del carrito actual.
// @Tags			Cart
// @Accept			json
// @Produce		json
// @Param			articleId		path		string					true	"ID de articlo"
// @Param			Authorization	header		string					true	"bearer {token}"
// @Success		200				{object}	cart.Cart				"Cart"
// @Param			body			body		cart.AddArticleData		true	"Articulo a Agregar"
// @Failure		400				{object}	apperr.ErrValidation	"Bad Request"
// @Failure		401				{object}	apperr.ErrCustom		"Unauthorized"
// @Failure		404				{object}	apperr.ErrCustom		"Not Found"
// @Failure		500				{object}	apperr.ErrCustom		"Internal Server Error"
// @Router			/v1/cart/article [post]
//
// Inicializa las rutas
func initPostCartArticle() {
	engine.Router().POST(
		"/v1/cart/article",
		middlewares.ValidateAuthentication,
		addArticle,
	)
}

func addArticle(c *gin.Context) {
	var options []interface{}
	if mocks, ok := c.Get("mocks"); ok {
		options = mocks.([]interface{})
	}

	user := c.MustGet("user").(security.User)
	body := cart.AddArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	_, err := service.AddArticle(user.ID, body, options...)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
