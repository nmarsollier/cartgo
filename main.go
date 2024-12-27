package main

import (
	"github.com/nmarsollier/cartgo/internal/engine/di"
	"github.com/nmarsollier/cartgo/internal/engine/env"
	"github.com/nmarsollier/cartgo/internal/engine/log"
	server "github.com/nmarsollier/cartgo/internal/graph"
	"github.com/nmarsollier/cartgo/internal/rabbit"
	"github.com/nmarsollier/cartgo/internal/rest"
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
	deps := di.NewInjector(log.Get(env.Get().FluentUrl))

	go server.Start(deps.Logger())

	rabbit.Init(deps)

	rest.Start()
}
