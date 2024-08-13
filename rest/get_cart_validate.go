package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// @Summary		Validar Carrito
// @Description	Valida el carrito para checkout
// @Tags			Cart
// @Accept			json
// @Produce		json
// @Param			Authorization	header	string	true	"bearer {token}"
// @Success		200				"No Content"
// @Failure		400				{object}	apperr.ValidationErr	"Bad Request"
// @Failure		401				{object}	engine.ErrorData		"Unauthorized"
// @Failure		404				{object}	engine.ErrorData		"Not Found"
// @Failure		500				{object}	engine.ErrorData		"Internal Server Error"
// @Router			/v1/cart/validate [get]
//
// Valida el carrito para checkout.
func initGetCartValidate() {
	engine.Router().GET(
		"/v1/cart/validate",
		engine.ValidateAuthentication,
		validate,
	)
}

func validate(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)

	ctx := engine.TestCtx(c)
	currentCart, err := cart.CurrentCart(user.ID, ctx...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	err = service.ValidateCheckout(currentCart, token, ctx...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
