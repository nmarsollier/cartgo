package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/internal/engine/log"
	"github.com/nmarsollier/cartgo/internal/engine/strs"
	"github.com/nmarsollier/cartgo/internal/security"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type LogoutConsumer interface {
	ConsumeLogout() error
}

func NewLogoutConsumer(fluentUrl string, rabbitUrl string, secService security.SecurityService) LogoutConsumer {
	logger := log.Get(fluentUrl).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "logout").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "auth").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &logoutConsumer{
		logger:     logger,
		rabbitUrl:  rabbitUrl,
		secService: secService,
	}
}

type logoutConsumer struct {
	logger     log.LogRusEntry
	rabbitUrl  string
	secService security.SecurityService
}

//	@Summary		Mensage Rabbit logout
//	@Description	Escucha de mensajes logout desde auth.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	logoutMessage	true	"Estructura general del mensage"
//	@Router			/rabbit/logout [get]
//
// Escucha de mensajes logout desde auth.
func (c *logoutConsumer) ConsumeLogout() error {
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
		"auth",   // name
		"fanout", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
		c.logger.Error(err)
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
		c.logger.Error(err)
		return err
	}

	go func() {
		for d := range mgs {
			newMessage := &logoutMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				c.logger.WithField(log.LOG_FIELD_CORRELATION_ID, getLogoutCorrelationId(newMessage))
				c.logger.Info("Incoming logout :", string(body))

				c.secService.Invalidate(newMessage.Message)
				c.logger.Info("Consumed :", strs.ToJson(newMessage))

			} else {
				c.logger.Error(err)
			}
		}
	}()

	c.logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

type logoutMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       string `json:"message" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"`
}

func getLogoutCorrelationId(c *logoutMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
