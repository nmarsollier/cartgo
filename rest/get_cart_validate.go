package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/services"
)

//	@Summary		Validar Carrito
//	@Description	Valida el carrito para checkout
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/validate [get]
//
// Valida el carrito para checkout.
func initGetCartValidate() {
	server.Router().GET(
		"/v1/cart/validate",
		server.ValidateAuthentication,
		validate,
	)
}

func validate(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)

	ctx := server.GinCtx(c)
	currentCart, err := cart.CurrentCart(user.ID, ctx...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	err = services.ValidateCheckout(currentCart, token, ctx...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
