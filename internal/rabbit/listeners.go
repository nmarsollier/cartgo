package rabbit

import (
	"github.com/nmarsollier/cartgo/internal/di"
)

func Init(di di.Injector) {
	logger := di.Logger()

	go listenLogout(logger)

	go listenArticleValidation(logger)

	go listenOrderPlaced(logger)

}
