package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
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
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/article/:articleId [delete]
//
// Elimina un articulo del carrito actual.
func initDeleteCart() {
	server.Router().DELETE(
		"/v1/cart/article/:articleId",
		server.ValidateAuthentication,
		deleteArticle,
	)
}

func deleteArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	deps := server.GinDeps(c)
	_, err := cart.RemoveArticle(user.ID, articleId, deps...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
