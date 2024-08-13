package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// Eliminar 1 al articulo actual.
// @Summary		Agregar 1 Articulo
// @Description	Eliminar 1 al articulo actual.
// @Tags			Cart
// @Accept			json
// @Produce		json
// @Param			articleId		path		string					true	"ID de articlo"
// @Param			Authorization	header		string					true	"bearer {token}"
// @Success		200				{object}	cart.Cart				"Cart"
// @Failure		400				{object}	apperr.ErrValidation	"Bad Request"
// @Failure		401				{object}	apperr.ErrCustom		"Unauthorized"
// @Failure		404				{object}	apperr.ErrCustom		"Not Found"
// @Failure		500				{object}	apperr.ErrCustom		"Internal Server Error"
// @Router			/v1/cart/article/:articleId/decrement [post]
//
// Init inicializa las rutas
func initPostCartArticleDecrement() {
	engine.Router().POST(
		"/v1/cart/article/:articleId/decrement",
		engine.ValidateAuthentication,
		decrementArticle,
	)
}

func decrementArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	article := cart.AddArticleData{
		ArticleId: articleId,
		Quantity:  -1,
	}

	ctx := engine.TestCtx(c)
	_, err := service.AddArticle(user.ID, article, ctx...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
