package rabbit

import (
	"time"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/di"
	"github.com/nmarsollier/cartgo/internal/env"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
)

//	@Summary		Mensage Rabbit article_exist/cart_article_exist
//	@Description	Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	rbt.InputMessage[cart.ValidationEvent]	true	"Mensaje"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func listenArticleValidation(logger log.LogRusEntry) {
	for {
		err := rbt.ConsumeRabbitEvent[*cart.ValidationEvent](
			env.Get().FluentURL,
			env.Get().RabbitURL,
			env.Get().ServerName,
			"article_exist",
			"direct",
			"cart_article_exist",
			"cart_article_exist",
			processArticleExist,
		)

		if err != nil {
			logger.Error(err)
		}
		logger.Info("RabbitMQ consumeCart conectando en 5 segundos.")
		time.Sleep(5 * time.Second)
	}
}

func processArticleExist(logger log.LogRusEntry, newMessage *rbt.InputMessage[*cart.ValidationEvent]) {
	deps := di.NewInjector(logger)

	data := newMessage.Message

	err := deps.CartService().ProcessArticleData(data)
	if err != nil {
		deps.Logger().Error(err)
		return
	}
}
