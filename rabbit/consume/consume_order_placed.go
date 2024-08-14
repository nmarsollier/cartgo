package consume

import (
	"encoding/json"
	"log"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit order/order-placed
//	@Description	Cuando se recibe order-placed se actualiza el order id del carrito. No se respode a este evento.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeOrderPlacedMessage	true	"Message para Type = order-placed"
//	@Router			/rabbit/order-placed [get]
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
		true,       // auto-ack
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

			glog.Info("Rabbit Consume: ", string(body))
			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "order-placed":
					processOrderPlaced(newMessage)
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

	log.Print("Order Placed processed : " + data.CartId)
}

type consumeOrderPlacedMessage struct {
	Type     string `json:"type" example:"order-placed" `
	Queue    string `json:"queue" example:"" `
	Exchange string `json:"exchange" example:"" `
	Message  *cart.OrderPlacedEvent
}
