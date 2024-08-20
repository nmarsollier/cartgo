package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/strs"
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
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"",             // routing key
		"order_placed", // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	glog.Info("RabbitMQ consumeOrderPlaced conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeOrderPlacedMessage{}
			body := d.Body

			glog.Info("Incomming order_placed :", string(body))
			err = json.Unmarshal(body, newMessage)
			if err == nil {
				processOrderPlaced(newMessage)

				if err := d.Ack(false); err != nil {
					glog.Info("Failed ACK order_placed : ", strs.ToJson(newMessage), err)
				} else {
					glog.Info("Consumed order_placed :", strs.ToJson(newMessage))
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processOrderPlaced(newMessage *consumeOrderPlacedMessage, ctx ...interface{}) {
	data := newMessage.Message

	err := cart.ProcessOrderPlaced(data, ctx...)
	if err != nil {
		glog.Error(err)
		return
	}
}

type consumeOrderPlacedMessage struct {
	Message *cart.OrderPlacedEvent
}
