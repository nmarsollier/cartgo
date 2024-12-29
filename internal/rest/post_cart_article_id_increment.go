package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
	"github.com/nmarsollier/commongo/security"
)

//	@Summary		Agregar 1 Articulo
//	@Description	Agregar 1 al articulo actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			articleId		path		string				true	"ID de articlo"
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	cart.Cart			"Cart"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/cart/article/:articleId/increment [post]
//
// Agregar 1 al articulo actual.
func initPostCartArticleIncrement(engine *gin.Engine) {
	engine.POST(
		"/cart/article/:articleId/increment",
		server.ValidateAuthentication,
		incrementArticle,
	)
}

func incrementArticle(c *gin.Context) {

	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	deps := server.GinDi(c)
	_, err := deps.Service().AddArticle(user.ID, articleId, 1)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
