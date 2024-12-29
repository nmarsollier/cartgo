package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
	"github.com/nmarsollier/commongo/security"
)

//	@Summary		Obtener carrito.
//	@Description	Obtiene el carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	cart.Cart			"Cart"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/cart [get]
//
// Obtiene el carrito actual del usuario.
func initGetCart(engine *gin.Engine) {
	engine.GET(
		"/cart",
		server.ValidateAuthentication,
		currentCart,
	)
}

func currentCart(c *gin.Context) {
	user := c.MustGet("user").(security.User)

	deps := server.GinDi(c)
	cart, err := deps.Service().GetCurrentCart(user.ID)

	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, cart)
}
