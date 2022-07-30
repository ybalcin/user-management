package err

// Error struct implements error interface and uses as base error for packages.
// Validations are slice of ValidationError.
// Code is http code
type Error struct {
	Err         string            `json:"error"`
	Code        int               `json:"-"`
	Validations []ValidationError `json:"validations,omitempty"`
}

// New initializes new Error
func New(err error, code int) *Error {
	e := &Error{
		Code:        code,
		Validations: []ValidationError{},
	}

	if err != nil {
		e.Err = err.Error()
	}

	return e
}

// Error returns error message
func (e *Error) Error() string {
	return e.Err
}

// AddValidationErr adds validation error to validations field
func (e *Error) AddValidationErr(err ValidationError) {
	e.Validations = append(e.Validations, err)
}
