package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/rest/middlewares"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// Obtiene el carrito actual del usuario.
//
//	@Summary		Obtener carrito.
//	@Description	Obtiene el carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Success		200				{object}		cart.Cart			"Cart"
//
//	@Failure		400				{object}	errors.ErrValidation	"Bad Request"
//	@Failure		401				{object}	errors.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	errors.ErrCustom		"Not Found"
//	@Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/cart [get]
func init() {
	engine.Router().GET(
		"/v1/cart",
		middlewares.ValidateAuthentication,
		currentCart,
	)
}

func currentCart(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	cart, err := service.GetCurrentCart(user.ID)

	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, cart)
}
