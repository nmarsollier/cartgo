package main

import (
	"github.com/nmarsollier/cartgo/internal/di"
	"github.com/nmarsollier/cartgo/internal/env"
	"github.com/nmarsollier/cartgo/internal/graph"
	"github.com/nmarsollier/cartgo/internal/rabbit"
	"github.com/nmarsollier/cartgo/internal/rest"
	"github.com/nmarsollier/commongo/log"
)

//	@title			CartGo
//	@version		1.0
//	@description	Microservicio de Carrito.
//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com
//
//	@host			localhost:3003
//	@BasePath		/v1
//
// Main Method
func main() {
	deps := di.NewInjector(log.Get(env.Get().FluentURL, env.Get().ServerName))

	go graph.Start(deps.Logger())

	rabbit.Init(deps)

	rest.Start()
}
