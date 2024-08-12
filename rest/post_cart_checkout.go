package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/rest/middlewares"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
)

// Hace checkout del carrito actual.
// @Summary		Checkout Articulo
// @Description	Hace checkout del carrito actual
// @Tags			Cart
// @Accept			json
// @Produce		json
// @Param			articleId		path	string	true	"ID de articlo"
// @Param			Authorization	header	string	true	"bearer {token}"
// @Success		200				"No Content"
// @Failure		400				{object}	apperr.ErrValidation	"Bad Request"
// @Failure		401				{object}	apperr.ErrCustom		"Unauthorized"
// @Failure		404				{object}	apperr.ErrCustom		"Not Found"
// @Failure		500				{object}	apperr.ErrCustom		"Internal Server Error"
// @Router			/v1/cart/checkouts [post]
//
// Inicializa las rutas
func init() {
	engine.Router().POST(
		"/v1/cart/checkout",
		middlewares.ValidateAuthentication,
		checkout,
	)
}

func checkout(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	token := c.MustGet("tokenString").(string)

	_, err := service.Checkout(user.ID, token)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
