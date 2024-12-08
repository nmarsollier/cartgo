// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nestor Marsollier",
            "email": "nmarsollier@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rabbit/article_exist": {
            "get": {
                "description": "Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit article_exist/cart_article_exist",
                "parameters": [
                    {
                        "description": "Mensaje",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.consumeArticleExistMessage"
                        }
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Solicitamos las validaciones ar articulos a catalogo. Responde en article_exist/cart_article_exist.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Emite Validar Artículos a Cart article_exist/article_exist",
                "parameters": [
                    {
                        "description": "Mensage de validacion article_exist/cart_article_exist",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/emit.SendValidationMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/logout": {
            "get": {
                "description": "Escucha de mensajes logout desde auth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit logout",
                "parameters": [
                    {
                        "description": "Estructura general del mensage",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.logoutMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/order_placed": {
            "get": {
                "description": "Cuando se recibe order_placed se actualiza el order id del carrito. No se respode a este evento.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit order_placed/order_placed",
                "parameters": [
                    {
                        "description": "Message order_placed",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.consumeOrderPlacedMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/place_order": {
            "put": {
                "description": "Cuando se hace checkout enviamos un comando a orders para que inicie el proceso de la orden.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Emite place_order/place_order",
                "parameters": [
                    {
                        "description": "Place order",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/emit.SendPlacedMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/cart": {
            "get": {
                "description": "Obtiene el carrito actual.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Obtener carrito.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart",
                        "schema": {
                            "$ref": "#/definitions/cart.Cart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/cart/article": {
            "post": {
                "description": "Agrega un articulo del carrito actual.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Agregar Articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Articulo a Agregar",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.AddArticleData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart",
                        "schema": {
                            "$ref": "#/definitions/cart.Cart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/cart/article/:articleId": {
            "delete": {
                "description": "Elimina un articulo del carrito actual.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Eliminar Articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/cart/article/:articleId/decrement": {
            "post": {
                "description": "Eliminar 1 al articulo actual.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Agregar 1 Articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart",
                        "schema": {
                            "$ref": "#/definitions/cart.Cart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/cart/article/:articleId/increment": {
            "post": {
                "description": "Agregar 1 al articulo actual.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Agregar 1 Articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart",
                        "schema": {
                            "$ref": "#/definitions/cart.Cart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/cart/checkouts": {
            "post": {
                "description": "Hace checkout del carrito actual",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Checkout Articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/cart/validate": {
            "get": {
                "description": "Valida el carrito para checkout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Validar Carrito",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cart.Article": {
            "type": "object",
            "required": [
                "articleId",
                "quantity"
            ],
            "properties": {
                "articleId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "quantity": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "valid": {
                    "type": "boolean"
                },
                "validated": {
                    "type": "boolean"
                }
            }
        },
        "cart.Cart": {
            "type": "object",
            "required": [
                "articles",
                "userId"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                },
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/cart.Article"
                    }
                },
                "created": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "orderId": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                },
                "userId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "userIdEnabled": {
                    "type": "string"
                }
            }
        },
        "cart.OrderPlacedEvent": {
            "type": "object",
            "properties": {
                "cartId": {
                    "type": "string",
                    "example": "CartId"
                },
                "orderId": {
                    "type": "string",
                    "example": "OrderId"
                },
                "valid": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "cart.ValidationEvent": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string",
                    "example": "ArticleId"
                },
                "referenceId": {
                    "type": "string",
                    "example": "UserId"
                },
                "valid": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "consume.consumeArticleExistMessage": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "message": {
                    "$ref": "#/definitions/cart.ValidationEvent"
                }
            }
        },
        "consume.consumeOrderPlacedMessage": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "message": {
                    "$ref": "#/definitions/cart.OrderPlacedEvent"
                }
            }
        },
        "consume.logoutMessage": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "message": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"
                }
            }
        },
        "emit.ArticleValidationData": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string",
                    "example": "ArticleId"
                },
                "referenceId": {
                    "type": "string",
                    "example": "UserId"
                }
            }
        },
        "emit.PlaceArticlesData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "ArticleId"
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "emit.PlacedData": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/emit.PlaceArticlesData"
                    }
                },
                "cartId": {
                    "type": "string",
                    "example": "CartId"
                },
                "userId": {
                    "type": "string",
                    "example": "UserId"
                }
            }
        },
        "emit.SendPlacedMessage": {
            "type": "object"
        },
        "emit.SendValidationMessage": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "exchange": {
                    "type": "string",
                    "example": "cart"
                },
                "message": {
                    "$ref": "#/definitions/emit.ArticleValidationData"
                },
                "routing_key": {
                    "type": "string",
                    "example": ""
                }
            }
        },
        "errs.ValidationErr": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/errs.errField"
                    }
                }
            }
        },
        "errs.errField": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "rest.AddArticleData": {
            "type": "object",
            "required": [
                "articleId",
                "quantity"
            ],
            "properties": {
                "articleId": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "quantity": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                }
            }
        },
        "server.ErrorData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3003",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "CartGo",
	Description:      "Microservicio de Carrito.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
