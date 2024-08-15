package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/errs"
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
	if err := validateToken(c); err != nil {
		c.Error(err)
		c.Abort()
		return
	}
}

// get token from Authorization header
func HeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}

func validateToken(c *gin.Context) error {
	tokenString, err := HeaderToken(c)
	if err != nil {
		glog.Error(err)
		return errs.Unauthorized
	}

	ctx := TestCtx(c)
	user, err := security.Validate(tokenString, ctx...)
	if err != nil {
		glog.Error(err)
		return errs.Unauthorized
	}

	c.Set("tokenString", tokenString)
	c.Set("user", *user)

	return nil
}
