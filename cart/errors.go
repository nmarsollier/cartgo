package cart

import "github.com/nmarsollier/cartgo/tools/apperr"

var ErrID = apperr.NewValidation().Add("id", "Invalid")
