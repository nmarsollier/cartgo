package consume

import (
	"time"

	"github.com/nmarsollier/cartgo/log"
)

func Init() {
	logger := log.Get().
		WithField("Controller", "Rabbit").
		WithField("Method", "Consume")

	go func() {
		for {
			err := consumeLogout()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeArticleExist()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeCart conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeOrderPlaced()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

}
