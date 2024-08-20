package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/strs"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit article_exist/cart_article_exist
//	@Description	Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeArticleDataMessage	true	"Mensaje"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func consumeArticleExist() error {
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
		"article_exist", // name
		"direct",        // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"cart_article_exist", // name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,           // queue name
		"cart_article_exist", // routing key
		"article_exist",      // exchange
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

	glog.Info("RabbitMQ consumeCart conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleDataMessage{}
			body := d.Body
			glog.Info("Incomming article_exist :", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				processArticleExist(newMessage)

				if err := d.Ack(false); err != nil {
					glog.Info("Failed ACK article_exist :", strs.ToJson(newMessage), err)
				} else {
					glog.Info("Consumed article_exist :", strs.ToJson(newMessage))
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processArticleExist(newMessage *consumeArticleDataMessage, ctx ...interface{}) {
	data := newMessage.Message

	err := cart.ProcessArticleData(data, ctx...)
	if err != nil {
		glog.Error(err)
		return
	}
}

type consumeArticleDataMessage struct {
	Message *cart.ValidationEvent
}
