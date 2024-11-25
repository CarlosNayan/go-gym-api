package errors

type InvalidCoordinatesError struct{}

func (e *InvalidCoordinatesError) Error() string {
	return "Invalid coordinates."
}

func (e *InvalidCoordinatesError) StatusCode() int {
	return 400
}
