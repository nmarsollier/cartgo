package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/services"
)

type AddArticleData struct {
	ArticleId string `bson:"articleId" validate:"required,min=1,max=100"`
	Quantity  int    `bson:"quantity" validate:"required,min=1,max=100"`
}

//	@Summary		Agregar Articulo
//	@Description	Agrega un articulo del carrito actual.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			articleId		path		string				true	"ID de articlo"
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	cart.Cart			"Cart"
//	@Param			body			body		AddArticleData		true	"Articulo a Agregar"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/cart/article [post]
//
// Agrega un articulo del carrito actual.
func initPostCartArticle() {
	server.Router().POST(
		"/v1/cart/article",
		server.ValidateAuthentication,
		addArticle,
	)
}

func addArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	body := AddArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	ctx := server.GinCtx(c)
	_, err := services.AddArticle(user.ID, body.ArticleId, body.Quantity, ctx...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
