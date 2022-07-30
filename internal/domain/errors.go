package domain

import "github.com/ybalcin/user-management/pkg/err"

var (
	UserNameCannotBeEmptyError     = err.NewValidationError("name", "user name cannot be empty")
	UserEmailCannotBeEmptyError    = err.NewValidationError("email", "user email cannot be empty")
	UserPasswordCannotBeEmptyError = err.NewValidationError("password", "user password cannot be empty")
	UserEmailInvalidError          = err.NewValidationError("email", "user email is invalid")
)
