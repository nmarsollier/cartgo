package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

//	@Summary		Checkout Articulo
//	@Description	Hace checkout del carrito actual
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			articleId		path	string	true	"ID de articlo"
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/checkouts [post]
//
// Hace checkout del carrito actual.
func initPostCartCheckout() {
	server.Router().POST(
		"/v1/cart/checkout",
		server.ValidateAuthentication,
		checkout,
	)
}

func checkout(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)

	ctx := server.GinCtx(c)
	_, err := service.Checkout(user.ID, token, ctx...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
