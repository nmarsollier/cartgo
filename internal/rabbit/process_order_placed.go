package rabbit

import (
	"time"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/di"
	"github.com/nmarsollier/cartgo/internal/env"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
)

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe order_placed se actualiza el order id del carrito. No se respode a este evento.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	rbt.InputMessage[cart.OrderPlacedEvent]	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed

func listenOrderPlaced(logger log.LogRusEntry) {
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
}

func processOrderPlaced(logger log.LogRusEntry, newMessage *rbt.InputMessage[*cart.OrderPlacedEvent]) {
	deps := di.NewInjector(logger)

	data := newMessage.Message

	err := deps.CartService().ProcessOrderPlaced(data)
	if err != nil {
		deps.Logger().Error(err)
		return
	}
}
