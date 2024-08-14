package cart

import "github.com/nmarsollier/cartgo/tools/errs"

var ErrID = errs.NewValidation().Add("id", "Invalid")
