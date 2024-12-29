package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
	"github.com/nmarsollier/commongo/security"
)

//	@Summary		Eliminar Articulo
//	@Description	Elimina un articulo del carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			articleId		path	string	true	"ID de articlo"
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/cart/article/:articleId [delete]
//
// Elimina un articulo del carrito actual.
func initDeleteCart(engine *gin.Engine) {
	engine.DELETE(
		"/cart/article/:articleId",
		server.ValidateAuthentication,
		deleteArticle,
	)
}

func deleteArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	deps := server.GinDi(c)
	_, err := deps.CartService().RemoveArticle(user.ID, articleId)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
