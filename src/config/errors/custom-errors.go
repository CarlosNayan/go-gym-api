package errors

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) StatusCode() int {
	return e.Code
}
