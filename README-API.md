# CartGo
Microservicio de Carrito.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/article-data

#### GET
##### Summary

Mensage Rabbit order/article-data

##### Description

Antes de iniciar las operaciones se validan los artículos contra el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| article-data | body | Message para Type = article-data | Yes | [r_consume.ConsumeArticleDataMessage](#r_consumeconsumearticledatamessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/cart/article-exist

#### PUT
##### Summary

Emite Validar Artículos a Cart cart/article-exist

##### Description

Antes de iniciar las operaciones se validan los artículos contra el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Mensage de validacion | Yes | [r_emit.SendValidationMessage](#r_emitsendvalidationmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/cart/place-order

#### PUT
##### Summary

Emite Placed Order desde Cart

##### Description

Emite Placed Order desde Cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Mensage de validacion | Yes | [r_emit.SendPlacedMessage](#r_emitsendplacedmessage) |

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
| body | body | Estructura general del mensage | Yes | [r_consume.LogoutMessage](#r_consumelogoutmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/order-placed

#### GET
##### Summary

Mensage Rabbit order/order-placed

##### Description

Antes de iniciar las operaciones se validan los artículos contra el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| article-data | body | Message para Type = article-data | Yes | [r_consume.ConsumeOrderPlacedMessage](#r_consumeconsumeorderplacedmessage) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

---
### Models

#### apperr.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [apperr.errField](#apperrerrfield) ] |  | No |

#### apperr.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

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
| cartId | string |  | No |
| orderId | string |  | No |
| valid | boolean |  | No |

#### cart.ValidationEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| referenceId | string |  | No |
| valid | boolean |  | No |

#### engine.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### r_consume.ConsumeArticleDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [cart.ValidationEvent](#cartvalidationevent) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### r_consume.ConsumeOrderPlacedMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [cart.OrderPlacedEvent](#cartorderplacedevent) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### r_consume.LogoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| type | string |  | No |

#### r_emit.ArticleValidationData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| referenceId | string |  | No |

#### r_emit.PlaceArticlesData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| quantity | integer |  | No |

#### r_emit.PlacedData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [r_emit.PlaceArticlesData](#r_emitplacearticlesdata) ] |  | No |
| cartId | string |  | No |
| userId | string |  | No |

#### r_emit.SendPlacedMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [r_emit.PlacedData](#r_emitplaceddata) |  | No |
| queue | string |  | No |
| type | string |  | No |

#### r_emit.SendValidationMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [r_emit.ArticleValidationData](#r_emitarticlevalidationdata) |  | No |
| queue | string |  | No |
| type | string |  | No |
