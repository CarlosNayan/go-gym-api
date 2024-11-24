package errors

type InvalidCoordinatesError struct{}

func (e *InvalidCoordinatesError) Error() string {
	return "Resource not found."
}

func (e *InvalidCoordinatesError) StatusCode() int {
	return 400
}
