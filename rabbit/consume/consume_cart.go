package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit order/article-exist
//	@Description	Luego de solicitar validaciones de catalogo, las validaciones las recibimos en esta Queue, con el mensaje type article-data.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeArticleDataMessage	true	"Message para Type = article-exist"
//	@Router			/rabbit/article-exist [get]
//
// Validar Artículos
func consumeOrders() error {
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
		"cart",   // name
		"direct", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"cart", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"cart",     // routing key
		"cart",     // exchange
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

	glog.Info("RabbitMQ consumeCart conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleDataMessage{}
			body := d.Body
			glog.Info(string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "article-exist":
					processArticleData(newMessage)
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processArticleData(newMessage *consumeArticleDataMessage, ctx ...interface{}) {
	data := newMessage.Message

	err := cart.ProcessArticleData(data, ctx...)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Info("Article exist completed : ", data)
}

type consumeArticleDataMessage struct {
	Type     string `json:"type" example:"article-exist"`
	Queue    string `json:"queue" example:""`
	Exchange string `json:"exchange" example:""`
	Message  *cart.ValidationEvent
}
