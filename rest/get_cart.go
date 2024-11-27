package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/services"
)

//	@Summary		Obtener carrito.
//	@Description	Obtiene el carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	cart.Cart			"Cart"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/cart [get]
//
// Obtiene el carrito actual del usuario.
func initGetCart() {
	server.Router().GET(
		"/v1/cart",
		server.ValidateAuthentication,
		currentCart,
	)
}

func currentCart(c *gin.Context) {
	user := c.MustGet("user").(security.User)

	deps := server.GinDeps(c)
	cart, err := services.GetCurrentCart(user.ID, deps...)

	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, cart)
}
