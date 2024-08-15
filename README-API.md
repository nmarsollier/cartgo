# CartGo
Microservicio de Carrito.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/article-exist

#### GET
##### Summary

Mensage Rabbit order/article-exist

##### Description

Luego de solicitar validaciones de catalogo, las validaciones las recibimos en esta Queue, con el mensaje type article-data.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| type | body | Message para Type = article-exist | Yes | [consume.consumeArticleDataMessage](#consumeconsumearticledatamessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/cart/article-exist

#### PUT
##### Summary

Emite Validar Artículos a Cart cart/article-exist

##### Description

Solicitamos las validaciones ar articulos a catalogo. Queue y Exchange es donde nos reponde.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Mensage de validacion | Yes | [emit.SendValidationMessage](#emitsendvalidationmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/cart/place-order

#### PUT
##### Summary

Emite Placed Order desde Cart

##### Description

Cuando se hace checkout enviamos un comando a orders para que inicie el proceso de la orden.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Place order | Yes | [emit.SendPlacedMessage](#emitsendplacedmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [consume.logoutMessage](#consumelogoutmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/order-placed

#### GET
##### Summary

Mensage Rabbit order/order-placed

##### Description

Cuando se recibe order-placed se actualiza el order id del carrito. No se respode a este evento.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| type | body | Message para Type = order-placed | Yes | [consume.consumeOrderPlacedMessage](#consumeconsumeorderplacedmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### /v1/cart

#### GET
##### Summary

Obtener carrito.

##### Description

Obtiene el carrito actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/cart/article

#### POST
##### Summary

Agregar Articulo

##### Description

Agrega un articulo del carrito actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | bearer {token} | Yes | string |
| body | body | Articulo a Agregar | Yes | [cart.AddArticleData](#cartaddarticledata) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/cart/article/:articleId

#### DELETE
##### Summary

Eliminar Articulo

##### Description

Elimina un articulo del carrito actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/cart/article/:articleId/decrement

#### POST
##### Summary

Agregar 1 Articulo

##### Description

Eliminar 1 al articulo actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/cart/article/:articleId/increment

#### POST
##### Summary

Agregar 1 Articulo

##### Description

Agregar 1 al articulo actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/cart/checkouts

#### POST
##### Summary

Checkout Articulo

##### Description

Hace checkout del carrito actual

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/cart/validate

#### GET
##### Summary

Validar Carrito

##### Description

Valida el carrito para checkout

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

---
### Models

#### cart.AddArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | Yes |
| quantity | integer |  | Yes |

#### cart.Article

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | Yes |
| quantity | integer |  | Yes |
| valid | boolean |  | No |
| validated | boolean |  | No |

#### cart.Cart

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _id | string |  | No |
| articles | [ [cart.Article](#cartarticle) ] |  | Yes |
| created | string |  | No |
| enabled | boolean |  | No |
| orderId | string |  | No |
| updated | string |  | No |
| userId | string |  | Yes |

#### cart.OrderPlacedEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cartId | string | *Example:* `"CartId"` | No |
| orderId | string | *Example:* `"OrderId"` | No |
| valid | boolean | *Example:* `true` | No |

#### cart.ValidationEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| referenceId | string | *Example:* `"UserId"` | No |
| valid | boolean | *Example:* `true` | No |

#### consume.consumeArticleDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string | *Example:* `""` | No |
| message | [cart.ValidationEvent](#cartvalidationevent) |  | No |
| queue | string | *Example:* `""` | No |
| type | string | *Example:* `"article-exist"` | No |

#### consume.consumeOrderPlacedMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string | *Example:* `""` | No |
| message | [cart.OrderPlacedEvent](#cartorderplacedevent) |  | No |
| queue | string | *Example:* `""` | No |
| type | string | *Example:* `"order-placed"` | No |

#### consume.logoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string | *Example:* `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"` | No |
| type | string | *Example:* `"logout"` | No |

#### emit.ArticleValidationData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| referenceId | string | *Example:* `"UserId"` | No |

#### emit.PlaceArticlesData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string | *Example:* `"ArticleId"` | No |
| quantity | integer | *Example:* `10` | No |

#### emit.PlacedData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [emit.PlaceArticlesData](#emitplacearticlesdata) ] |  | No |
| cartId | string | *Example:* `"CartId"` | No |
| userId | string | *Example:* `"UserId"` | No |

#### emit.SendPlacedMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| emit.SendPlacedMessage | object |  |  |

#### emit.SendValidationMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string | *Example:* `"cart"` | No |
| message | [emit.ArticleValidationData](#emitarticlevalidationdata) |  | No |
| queue | string | *Example:* `"cart"` | No |
| type | string | *Example:* `"article-exist"` | No |

#### errs.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errs.errField](#errserrfield) ] |  | No |

#### errs.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### server.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
