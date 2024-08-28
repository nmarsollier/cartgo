package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/log"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/strs"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe order_placed se actualiza el order id del carrito. No se respode a este evento.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeOrderPlacedMessage	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed
func consumeOrderPlaced() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "cart_order_placed").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"",             // routing key
		"order_placed", // exchange
		false,
		nil)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	go func() {
		for d := range mgs {
			newMessage := &consumeOrderPlacedMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := logger.WithField(log.LOG_FIELD_CORRELATION_ID, getOrderPlacedCorrelationId(newMessage))
				l.Info("Incoming order_placed :", string(body))

				processOrderPlaced(newMessage, l)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK order_placed : ", strs.ToJson(newMessage), err)
				} else {
					l.Info("Consumed order_placed :", strs.ToJson(newMessage))
				}
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processOrderPlaced(newMessage *consumeOrderPlacedMessage, ctx ...interface{}) {
	data := newMessage.Message

	err := cart.ProcessOrderPlaced(data, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
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
