package application

import "errors"

var (
	UserExistWithThatEmailError   = errors.New("user with that email already exists")
	IdCannotBeEmptyError          = errors.New("id cannot be empty")
	UserIsNotExistWithThatIdError = errors.New("user with that id does not exist")
)
