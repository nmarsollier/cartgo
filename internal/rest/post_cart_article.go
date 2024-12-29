package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
	"github.com/nmarsollier/commongo/security"
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
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/cart/article [post]
//
// Agrega un articulo del carrito actual.
func initPostCartArticle(engine *gin.Engine) {
	engine.POST(
		"/cart/article",
		server.ValidateAuthentication,
		addArticle,
	)
}

func addArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	body := AddArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		rst.AbortWithError(c, err)
		return
	}

	deps := server.GinDi(c)
	_, err := deps.Service().AddArticle(user.ID, body.ArticleId, body.Quantity)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
