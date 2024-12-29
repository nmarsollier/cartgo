package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
	"github.com/nmarsollier/commongo/security"
)

//	@Summary		Validar Carrito
//	@Description	Valida el carrito para checkout
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/cart/validate [get]
//
// Valida el carrito para checkout.
func initGetCartValidate(engine *gin.Engine) {
	engine.GET(
		"/cart/validate",
		server.ValidateAuthentication,
		validate,
	)
}

func validate(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)

	deps := server.GinDi(c)
	currentCart, err := deps.CartService().CurrentCart(user.ID)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	err = deps.Service().ValidateCheckout(currentCart, token)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
