package err

// ValidationError struct uses for validation error messages
type ValidationError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

// NewValidationError initializes validation error
func NewValidationError(key, message string) ValidationError {
	return ValidationError{
		Key:     key,
		Message: message,
	}
}

func (e *ValidationError) Error() string {
	return e.Message
}
