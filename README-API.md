# CartGo
Microservicio de Carrito.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /cart

#### GET
##### Summary

Obtener carrito.

##### Description

Obtiene el carrito actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /cart/article

#### POST
##### Summary

Agregar Articulo

##### Description

Agrega un articulo del carrito actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | Bearer {token} | Yes | string |
| body | body | Articulo a Agregar | Yes | [rest.AddArticleData](#restaddarticledata) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /cart/article/:articleId

#### DELETE
##### Summary

Eliminar Articulo

##### Description

Elimina un articulo del carrito actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | Bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /cart/article/:articleId/decrement

#### POST
##### Summary

Agregar 1 Articulo

##### Description

Eliminar 1 al articulo actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | Bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /cart/article/:articleId/increment

#### POST
##### Summary

Agregar 1 Articulo

##### Description

Agregar 1 al articulo actual.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | Bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Cart | [cart.Cart](#cartcart) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /cart/checkouts

#### POST
##### Summary

Checkout Articulo

##### Description

Hace checkout del carrito actual

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| articleId | path | ID de articlo | Yes | string |
| Authorization | header | Bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /cart/validate

#### GET
##### Summary

Validar Carrito

##### Description

Valida el carrito para checkout

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

---
### /rabbit/article_exist

#### GET
##### Summary

Mensage Rabbit article_exist/cart_article_exist

##### Description

Luego de solicitar validaciones de catalogo, Escucha article_exist/cart_article_exist.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| type | body | Mensaje | Yes | [rbt.InputMessage-cart_ValidationEvent](#rbtinputmessage-cart_validationevent) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit logout

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [rbt.InputMessage-string](#rbtinputmessage-string) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### Models

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

#### cart.ValidationEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| referenceId | string | *Example:* `"UserId"` | No |
| valid | boolean | *Example:* `true` | No |

#### errs.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errs.errField](#errserrfield) ] |  | No |

#### errs.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### rbt.InputMessage-cart_ValidationEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| correlation_id | string | *Example:* `"123123"` | No |
| message | [cart.ValidationEvent](#cartvalidationevent) |  | No |

#### rbt.InputMessage-string

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| correlation_id | string | *Example:* `"123123"` | No |
| message | string |  | No |

#### rest.AddArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | Yes |
| quantity | integer |  | Yes |

#### rst.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
