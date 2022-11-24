package originalErrors

import "fmt"

type ErrorType int

const (
	ErrorType0 ErrorType = 0
	ErrorType1 ErrorType = 1
	ErrorType2 ErrorType = 2
)

func ErrorWithType(e *errorString) string {
	return fmt.Sprintf("Type: %v, Error:%v", e.s, e.t.GetTypeName())
}

//func ErrorWithType1(e error) string {
//	var err errorString = e.(errorString)
//	return fmt.Sprintf("Type: %v, Error:%v", err.s, err.t.GetTypeName())
//}

func (e ErrorType) GetTypeName() string {
	switch e {
	case ErrorType0:
		return "ErrorType0"
	case ErrorType1:
		return "ErrorType1"
	case ErrorType2:
		return "ErrorType2"
	default:
		return "undefied errorType"
	}
}
