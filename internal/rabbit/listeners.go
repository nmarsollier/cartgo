package rabbit

import (
	"time"

	"github.com/nmarsollier/cartgo/internal/di"
)

func Init(di di.Injector) {
	logger := di.Logger()

	go func() {
		for {
			err := di.LogoutConsumer().ConsumeLogout()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := di.ArticleExistConsumer().ConsumeArticleExist()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeCart conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := di.ConsumeOrderPlaced().ConsumeOrderPlaced()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

}
