package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/engine"
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
func initGetCart() {
	engine.Router().GET(
		"/v1/cart",
		engine.ValidateAuthentication,
		currentCart,
	)
}

func currentCart(c *gin.Context) {
	user := c.MustGet("user").(security.User)

	ctx := engine.TestCtx(c)
	cart, err := service.GetCurrentCart(user.ID, ctx...)

	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, cart)
}
