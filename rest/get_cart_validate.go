package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/rest/middlewares"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// Valida el carrito para checkout.
//
//	@Summary		Validar Carrito
//	@Description	Valida el carrito para checkout
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Success		200				"No Content"
//
//	@Failure		400				{object}	errors.ErrValidation	"Bad Request"
//	@Failure		401				{object}	errors.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	errors.ErrCustom		"Not Found"
//	@Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/cart/validate [get]
func init() {
	engine.Router().GET(
		"/v1/cart/validate",
		middlewares.ValidateAuthentication,
		validate,
	)
}

func validate(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)
	currentCart, err := cart.CurrentCart(user.ID)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	err = service.ValidateCheckout(currentCart, token)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
