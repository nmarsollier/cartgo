package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/services"
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
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/article/:articleId/increment [post]
//
// Agregar 1 al articulo actual.
func initPostCartArticleIncrement() {
	server.Router().POST(
		"/v1/cart/article/:articleId/increment",
		server.ValidateAuthentication,
		incrementArticle,
	)
}

func incrementArticle(c *gin.Context) {

	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	deps := server.GinDeps(c)
	_, err := services.AddArticle(user.ID, articleId, 1, deps...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
