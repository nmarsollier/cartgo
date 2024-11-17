<!-- cSpell:language es -->

# Microservicio de Carrito

Este microservicio gestiona los artículos en el carrito por usuario.

Sólo existe un carrito activo por usuario.

Una vez hecha la compra el carrito se cierra y se genera una orden.

[Documentación de API](./README-API.md)

La documentación de las api también se pueden consultar desde el home del microservicio
que una vez levantado el servidor se puede navegar en [localhost:3003](http://localhost:3003/)

El servidor GraphQL puede navegar en [localhost:4003](http://localhost:4003/)

## Directorios

- **cart:** Logica de negocio del agregado carrito
- **services:** Domain services.
- **security:** Validaciones de usuario contra el MS de Auth
- **graph:** Servidor y Controllers GraphQL federation server
- **rabbit:** Servidor y Controllers RabbitMQ
- **rest:** Servidor y Controllers Rest
- **tools:** Herramientas varias

## Dependencias

### Auth

Las imágenes solo pueden subirse y descargarse por usuarios autenticados, ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

### Catalog

Los artículos en el carrito se validan contra el catalogo, ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

### Node 20

Seguir los pasos de instalación del sitio oficial

[nodejs.org](https://nodejs.org/en/)

### MongoDb

Ver tutorial de instalación en [ecommerce](https://github.com/nmarsollier/ecommerce).

### RabbitMQ

La comunicación con Catalog y Auth es a través de rabbit.

Ver tutorial de instalación en [ecommerce](https://github.com/nmarsollier/ecommerce).

## Ejecución

```bash
git config core.hooksPath .githooks
go install github.com/swaggo/gin-swagger/swaggerFiles
go install github.com/swaggo/gin-swagger
go install github.com/swaggo/swag/cmd/swag
go install github.com/golang/mock/mockgen@v1.6.0
go install github.com/99designs/gqlgen@v0.17.56
```

## Build y ejecución

```bash
go install
cartgo
```

## Swagger

Usamos [swaggo](https://github.com/swaggo/swag)

Requisitos

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

La documentacion la generamos con el comando

```bash
swag init
```

Para generar el archivo README-API.md

Requisito

```bash
sudo npm install -g swagger-markdown
```

y ejecutamos

```bash
npx swagger-markdown -i ./docs/swagger.yaml -o README-API.md
```

## Configuración del servidor

Este servidor se configura con variables de entorno

MONGO_URL=mongodb://host.docker.internal:27017
RABBIT_URL=amqp://host.docker.internal
AUTH_SERVICE_URL=http://host.docker.internal:3000
CATALOG_SERVICE_URL=http://host.docker.internal:3002

## Docker

Estos comandos son para dockerizar el microservicio desde el codigo descargado localmente.

### Build

```bash
docker build -t dev-cart-go .
```

### El contenedor

Mac | Windows

```bash
docker run -it --name dev-cart-go -p 3003:3003 -p 4003:4003 -v $PWD:/go/src/github.com/nmarsollier/cartgo dev-cart-go
```

Linux

```bash
docker run -it --add-host host.docker.internal:172.17.0.1 --name dev-cart-go -p 3003:3003 -p 4003:4003 -v $PWD:/go/src/github.com/nmarsollier/cartgo dev-cart-go
```
