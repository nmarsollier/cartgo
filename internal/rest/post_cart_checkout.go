package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
	"github.com/nmarsollier/commongo/security"
)

//	@Summary		Checkout Articulo
//	@Description	Hace checkout del carrito actual
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
//	@Router			/cart/checkouts [post]
//
// Hace checkout del carrito actual.
func initPostCartCheckout(engine *gin.Engine) {
	engine.POST(
		"/cart/checkout",
		server.ValidateAuthentication,
		checkout,
	)
}

func checkout(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)

	deps := server.GinDi(c)
	_, err := deps.Service().Checkout(user.ID, token)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
