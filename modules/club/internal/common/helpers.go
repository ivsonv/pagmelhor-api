package common

// func BindAndValidate(c echo.Context, item any) error {
// 	if err := c.Bind(item); err != nil {
// 		return NewError(ErrBindError, err.Error(), http.StatusBadRequest)
// 	}

// 	if err := c.Validate(item); err != nil {
// 		return NewError(ErrValidateError, err.Error(), http.StatusBadRequest)
// 	}

// 	return nil
// }
