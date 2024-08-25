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

//	@Summary		Mensage Rabbit article_exist/cart_article_exist
//	@Description	Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeArticleExistMessage	true	"Mensaje"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func consumeArticleExist() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "cart_article_exist").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
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
		"article_exist", // name
		"direct",        // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,           // queue name
		"cart_article_exist", // routing key
		"article_exist",      // exchange
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

	logger.Info("RabbitMQ consumeCart conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleExistMessage{}
			body := d.Body
			logger.Info(string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := logger.WithField(log.LOG_FIELD_CORRELATION_ID, getArticleExistCorrelationId(newMessage))

				processArticleExist(newMessage, l)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK article_exist :", strs.ToJson(newMessage), err)
				} else {
					l.Info("Consumed article_exist :", strs.ToJson(newMessage))
				}
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processArticleExist(newMessage *consumeArticleExistMessage, ctx ...interface{}) {
	data := newMessage.Message

	err := cart.ProcessArticleData(data, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return
	}
}

type consumeArticleExistMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       *cart.ValidationEvent
}

func getArticleExistCorrelationId(c *consumeArticleExistMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
