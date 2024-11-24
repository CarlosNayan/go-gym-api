package errors

type InvalidRequestBodyError struct{}

func (e *InvalidRequestBodyError) Error() string {
	return "Invalid request body."
}

func (e *InvalidRequestBodyError) StatusCode() int {
	return 400
}
