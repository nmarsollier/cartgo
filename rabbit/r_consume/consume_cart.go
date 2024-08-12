package r_consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

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
			newMessage := &ConsumeMessage{}
			body := d.Body
			glog.Info(string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "article-exist":
					articleMessage := &ConsumeArticleDataMessage{}
					if err := json.Unmarshal(body, articleMessage); err != nil {
						glog.Error(err)
						return
					}

					processArticleData(articleMessage)
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

// Validar Artículos
//
//	@Summary		Mensage Rabbit order/article-data
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article-data	body	ConsumeArticleDataMessage	true	"Message para Type = article-data"
//
//	@Router			/rabbit/article-data [get]
func processArticleData(newMessage *ConsumeArticleDataMessage, options ...interface{}) {
	data := newMessage.Message

	err := cart.ProcessArticleData(data, options...)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Info("Article exist completed : ", data)
}

type ConsumeArticleDataMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *cart.ValidationEvent
}
