package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// @Summary		Obtener carrito.
// @Description	Obtiene el carrito actual.
// @Tags			Cart
// @Accept			json
// @Produce		json
// @Param			Authorization	header		string					true	"bearer {token}"
// @Success		200				{object}		cart.Cart			"Cart"
// @Failure		400				{object}	apperr.ValidationErr	"Bad Request"
// @Failure		401				{object}	engine.ErrorData		"Unauthorized"
// @Failure		404				{object}	engine.ErrorData		"Not Found"
// @Failure		500				{object}	engine.ErrorData		"Internal Server Error"
// @Router			/v1/cart [get]
//
// Obtiene el carrito actual del usuario.
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
