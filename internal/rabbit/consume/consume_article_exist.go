package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/strs"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type ArticleExistConsumer interface {
	ConsumeArticleExist() error
}

func NewArticleExistConsumer(fluentUrl string, rabbitUrl string, service cart.CartService) ArticleExistConsumer {
	log := log.Get(fluentUrl, "cartgo").
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "cart_article_exist").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &articleExistConsumer{
		log:       log,
		rabbitUrl: rabbitUrl,
		service:   service,
	}
}

type articleExistConsumer struct {
	log       log.LogRusEntry
	rabbitUrl string
	service   cart.CartService
}

//	@Summary		Mensage Rabbit article_exist/cart_article_exist
//	@Description	Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			type	body	consumeArticleExistMessage	true	"Mensaje"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func (c *articleExistConsumer) ConsumeArticleExist() error {

	conn, err := amqp.Dial(c.rabbitUrl)
	if err != nil {
		c.log.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		c.log.Error(err)
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
		c.log.Error(err)
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
		c.log.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,           // queue name
		"cart_article_exist", // routing key
		"article_exist",      // exchange
		false,
		nil)
	if err != nil {
		c.log.Error(err)
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
		c.log.Error(err)
		return err
	}

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleExistMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				c.log.WithField(log.LOG_FIELD_CORRELATION_ID, getArticleExistCorrelationId(newMessage))
				c.log.Info("Incoming article_exist :", string(body))

				c.processArticleExist(newMessage)

				if err := d.Ack(false); err != nil {
					c.log.Info("Failed ACK article_exist :", strs.ToJson(newMessage), err)
				} else {
					c.log.Info("Consumed article_exist :", strs.ToJson(newMessage))
				}
			} else {
				c.log.Error(err)
			}
		}
	}()

	c.log.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func (c *articleExistConsumer) processArticleExist(newMessage *consumeArticleExistMessage) {
	data := newMessage.Message

	err := c.service.ProcessArticleData(data)
	if err != nil {
		c.log.Error(err)
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
