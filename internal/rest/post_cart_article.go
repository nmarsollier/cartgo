package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/internal/rest/engine"
	"github.com/nmarsollier/cartgo/internal/security"
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
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/cart/article [post]
//
// Agrega un articulo del carrito actual.
func initPostCartArticle() {
	engine.Router().POST(
		"/cart/article",
		engine.ValidateAuthentication,
		addArticle,
	)
}

func addArticle(c *gin.Context) {
	user := c.MustGet("user").(security.User)
	body := AddArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		engine.AbortWithError(c, err)
		return
	}

	deps := engine.GinDi(c)
	_, err := deps.Service().AddArticle(user.ID, body.ArticleId, body.Quantity)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
