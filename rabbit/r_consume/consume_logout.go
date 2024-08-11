package r_consume

import (
	"encoding/json"
	"fmt"

	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

type LogoutMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Escucha de mensajes logout desde auth.
//
//	@Summary		Mensage Rabbit
//	@Description	Escucha de mensajes logout desde auth.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	LogoutMessage	true	"Estructura general del mensage"
//	@Router			/rabbit/logout [get]
func consumeLogout() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
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
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
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
		return err
	}

	fmt.Println("RabbitMQ listenLogout conectado")

	go func() {
		for d := range mgs {
			newMessage := &LogoutMessage{}
			body := d.Body
			fmt.Println(string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				if newMessage.Type == "logout" {
					security.Invalidate(newMessage.Message)
				}
			}
		}
	}()

	fmt.Print("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}
