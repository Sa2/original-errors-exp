package originalErrors

import "fmt"

type errorString struct {
	s string
	t ErrorType // t is errorType
}

func New(text string, errorType ErrorType) error {
	return &errorString{
		s: text,
		t: errorType,
	}
}

func (e *errorString) Error() string {
	return fmt.Sprintf("Error: %v, Type: %v", e.s, e.t.GetTypeName())
}

func (e *errorString) ErrorType() ErrorType {
	return e.t
}
