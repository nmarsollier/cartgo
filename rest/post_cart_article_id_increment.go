package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

//	@Summary		Agregar 1 Articulo
//	@Description	Agregar 1 al articulo actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			articleId		path		string				true	"ID de articlo"
//	@Param			Authorization	header		string				true	"bearer {token}"
//	@Success		200				{object}	cart.Cart			"Cart"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/article/:articleId/increment [post]
//
// Agregar 1 al articulo actual.
func initPostCartArticleIncrement() {
	engine.Router().POST(
		"/v1/cart/article/:articleId/increment",
		engine.ValidateAuthentication,
		incrementArticle,
	)
}

func incrementArticle(c *gin.Context) {

	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	article := cart.AddArticleData{
		ArticleId: articleId,
		Quantity:  1,
	}

	ctx := engine.TestCtx(c)
	_, err := service.AddArticle(user.ID, article, ctx...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
