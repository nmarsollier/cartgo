package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/engine"
	"github.com/nmarsollier/cartgo/internal/security"
)

//	@Summary		Obtener carrito.
//	@Description	Obtiene el carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	cart.Cart			"Cart"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/cart [get]
//
// Obtiene el carrito actual del usuario.
func initGetCart() {
	engine.Router().GET(
		"/cart",
		engine.ValidateAuthentication,
		currentCart,
	)
}

func currentCart(c *gin.Context) {
	user := c.MustGet("user").(security.User)

	deps := engine.GinDi(c)
	cart, err := deps.Service().GetCurrentCart(user.ID)

	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, cart)
}
