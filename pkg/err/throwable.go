package err

import (
	"errors"
	"net/http"
)

// ThrowValidationError throws Error with validations and code http.StatusBadRequest
func ThrowValidationError(validations ...ValidationError) *Error {
	e := New(errors.New(SomeValidationErrorsOccurredText), http.StatusBadRequest)

	for _, v := range validations {
		e.Validations = append(e.Validations, v)
	}

	return e
}

// ThrowBadRequestError throws Error with code http.StatusBadRequest
func ThrowBadRequestError(e error) *Error {
	return New(e, http.StatusBadRequest)
}

// ThrowForbiddenError throws Error with code http.StatusForbidden
func ThrowForbiddenError(e error) *Error {
	return New(e, http.StatusForbidden)
}

// ThrowBadRequestErrorWithMessage throws Error with message and code http.StatusBadRequest
func ThrowBadRequestErrorWithMessage(message string) *Error {
	return New(errors.New(message), http.StatusBadRequest)
}

// ThrowInternalServerError throws Error with code http.StatusInternalServerError
func ThrowInternalServerError(e error) *Error {
	if e == nil {
		e = errors.New(AnErrorOccurredText)
	}

	return New(e, http.StatusInternalServerError)
}

// ThrowNotFoundError throws Error with http.StatusNotFound
func ThrowNotFoundError(e error) *Error {
	return New(e, http.StatusNotFound)
}
