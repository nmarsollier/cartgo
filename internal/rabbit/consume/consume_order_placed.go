package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/engine/log"
	"github.com/nmarsollier/cartgo/internal/engine/strs"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type OrderPlacedConsumer interface {
	ConsumeOrderPlaced() error
}

func NewOrderPlacedConsumer(fluentUrl string, rabbitUrl string, cart cart.CartService) OrderPlacedConsumer {
	logger := log.Get(fluentUrl).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "cart_order_placed").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &orderPlacedConsumer{
		logger:    logger,
		rabbitUrl: rabbitUrl,
		cart:      cart,
	}
}

type orderPlacedConsumer struct {
	logger    log.LogRusEntry
	rabbitUrl string
	cart      cart.CartService
}

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe order_placed se actualiza el order id del carrito. No se respode a este evento.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeOrderPlacedMessage	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed
func (c *orderPlacedConsumer) ConsumeOrderPlaced() error {
	conn, err := amqp.Dial(c.rabbitUrl)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		c.logger.Error(err)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"order_placed", // name
		"fanout",       // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"cart_order_placed", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"",             // routing key
		"order_placed", // exchange
		false,
		nil)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	go func() {
		for d := range mgs {
			newMessage := &consumeOrderPlacedMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				c.logger.WithField(log.LOG_FIELD_CORRELATION_ID, getOrderPlacedCorrelationId(newMessage))
				c.logger.Info("Incoming order_placed :", string(body))

				c.processOrderPlaced(newMessage)

				if err := d.Ack(false); err != nil {
					c.logger.Info("Failed ACK order_placed : ", strs.ToJson(newMessage), err)
				} else {
					c.logger.Info("Consumed order_placed :", strs.ToJson(newMessage))
				}
			} else {
				c.logger.Error(err)
			}
		}
	}()

	c.logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func (c *orderPlacedConsumer) processOrderPlaced(newMessage *consumeOrderPlacedMessage) {
	data := newMessage.Message

	err := c.cart.ProcessOrderPlaced(data)
	if err != nil {
		c.logger.Error(err)
		return
	}
}

type consumeOrderPlacedMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       *cart.OrderPlacedEvent
}

func getOrderPlacedCorrelationId(c *consumeOrderPlacedMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
