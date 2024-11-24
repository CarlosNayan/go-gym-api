package errors

type ResourceNotFoundError struct{}

func (e *ResourceNotFoundError) Error() string {
	return "Resource not found."
}

func (e *ResourceNotFoundError) StatusCode() int {
	return 404
}
