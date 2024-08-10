package rest

import (
	"fmt"

	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/tools/env"
)

// Start this server
func Start() {
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}
