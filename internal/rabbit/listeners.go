package rabbit

import (
	"time"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/di"
	"github.com/nmarsollier/cartgo/internal/env"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
)

func Init(di di.Injector) {
	logger := di.Logger()

	go func() {
		for {
			err := rbt.ConsumeRabbitEvent[string](
				env.Get().FluentURL,
				env.Get().RabbitURL,
				env.Get().ServerName,
				"auth",
				"fanout",
				"",
				"",
				processLogout,
			)

			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
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
	}()

	go func() {
		for {
			err := rbt.ConsumeRabbitEvent[*cart.OrderPlacedEvent](
				env.Get().FluentURL,
				env.Get().RabbitURL,
				env.Get().ServerName,
				"order_placed",
				"fanout",
				"cart_order_placed",
				"",
				processOrderPlaced,
			)
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

}

//	@Summary		Mensage Rabbit article_exist/cart_article_exist
//	@Description	Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	rbt.InputMessage[cart.ValidationEvent]	true	"Mensaje"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func processArticleExist(logger log.LogRusEntry, newMessage *rbt.InputMessage[*cart.ValidationEvent]) {
	deps := di.NewInjector(logger)

	data := newMessage.Message

	err := deps.CartService().ProcessArticleData(data)
	if err != nil {
		deps.Logger().Error(err)
		return
	}
}

//	@Summary		Mensage Rabbit logout
//	@Description	Escucha de mensajes logout desde auth.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	rbt.InputMessage[string]	true	"Estructura general del mensage"
//	@Router			/rabbit/logout [get]
//
// Escucha de mensajes logout desde auth.
func processLogout(logger log.LogRusEntry, newMessage *rbt.InputMessage[string]) {
	deps := di.NewInjector(logger)

	deps.SecurityService().Invalidate(newMessage.Message)
}

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe order_placed se actualiza el order id del carrito. No se respode a este evento.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	rbt.InputMessage[cart.OrderPlacedEvent]	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed
func processOrderPlaced(logger log.LogRusEntry, newMessage *rbt.InputMessage[*cart.OrderPlacedEvent]) {
	deps := di.NewInjector(logger)

	data := newMessage.Message

	err := deps.CartService().ProcessOrderPlaced(data)
	if err != nil {
		deps.Logger().Error(err)
		return
	}
}
