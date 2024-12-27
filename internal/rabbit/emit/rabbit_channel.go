package emit

import (
	"errors"

	"github.com/nmarsollier/cartgo/internal/engine/log"
	"github.com/streadway/amqp"
)

type RabbitChannel interface {
	ExchangeDeclare(name string, chType string) error
	Publish(exchange string, routingKey string, body []byte) error
}

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func NewChannel(
	rabbitUrl string,
	log log.LogRusEntry,
) (RabbitChannel, error) {
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return rabbitChannel{ch: channel}, nil
}

type rabbitChannel struct {
	ch *amqp.Channel
}

func (c rabbitChannel) ExchangeDeclare(
	name string,
	chType string,
) error {
	return c.ch.ExchangeDeclare(
		name,   // name
		chType, // type
		false,  // durable
		false,  // auto-deleted
		false,  // internal
		false,  // no-wait
		nil,    // arguments
	)
}
func (c rabbitChannel) Publish(
	exchange string,
	routingKey string,
	body []byte,
) error {
	return c.ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Body: body,
		})
}
