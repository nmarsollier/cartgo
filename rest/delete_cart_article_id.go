package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
)

//	@Summary		Eliminar Articulo
//	@Description	Elimina un articulo del carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			articleId		path	string	true	"ID de articlo"
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/article/:articleId [delete]
//
// Elimina un articulo del carrito actual.
func initDeleteCart() {
	engine.Router().DELETE(
		"/v1/cart/article/:articleId",
		engine.ValidateAuthentication,
		deleteArticle,
	)
}

func deleteArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	articleId := c.Param("articleId")

	ctx := engine.TestCtx(c)
	_, err := cart.RemoveArticle(user.ID, articleId, ctx...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
