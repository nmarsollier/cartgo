package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/log"
)

/**
 * @apiDefine AuthHeader
 *
 * @apiExample {String} Header Autorización
 *    Authorization=bearer {token}
 *
 * @apiErrorExample 401 Unauthorized
 *    HTTP/1.1 401 Unauthorized
 */

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	user, err := validateToken(c)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	deps := GinDeps(c)
	c.Set("logger", log.Get(deps...).WithField(log.LOG_FIELD_USER_ID, user.ID))
}

// get token from Authorization header
func HeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(strings.ToUpper(tokenString), "BEARER ") != 0 {
		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}

func validateToken(c *gin.Context) (*security.User, error) {
	tokenString, err := HeaderToken(c)
	if err != nil {
		return nil, errs.Unauthorized
	}

	deps := GinDeps(c)
	user, err := security.Validate(tokenString, deps...)
	if err != nil {
		return nil, errs.Unauthorized
	}

	c.Set("tokenString", tokenString)
	c.Set("user", *user)

	return user, nil
}
