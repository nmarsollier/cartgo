package r_consume

import (
	"encoding/json"
	"log"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

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
		queue.Name,          // queue name
		"cart_order_placed", // routing key
		"order_placed",      // exchange
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
			newMessage := &ConsumeMessage{}
			body := d.Body

			glog.Info("Rabbit Consume: ", string(body))
			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "order-placed":
					articleMessage := &ConsumeOrderPlacedMessage{}
					if err := json.Unmarshal(body, articleMessage); err != nil {
						glog.Error(err)
						return
					}

					processOrderPlaced(articleMessage)
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

// Consume Order Placed
//
//	@Summary		Mensage Rabbit order/order-placed
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article-data	body	ConsumeOrderPlacedMessage	true	"Message para Type = article-data"
//
//	@Router			/rabbit/order-placed [get]
func processOrderPlaced(newMessage *ConsumeOrderPlacedMessage) {
	data := newMessage.Message

	err := cart.ProcessOrderPlaced(data)
	if err != nil {
		glog.Error(err)
		return
	}

	log.Print("Order Placed processed : " + data.CartId)
}

type ConsumeOrderPlacedMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *cart.OrderPlacedEvent
}
