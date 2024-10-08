package main

import (
	"github.com/nmarsollier/cartgo/rabbit/consume"
	routes "github.com/nmarsollier/cartgo/rest"
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
	consume.Init()
	routes.Start()
}
