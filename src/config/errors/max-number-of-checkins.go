package errors

type MaxNumberOfCheckinsError struct{}

func (e *MaxNumberOfCheckinsError) Error() string {
	return "You have reached the maximum number of checkins."
}

func (e *MaxNumberOfCheckinsError) StatusCode() int {
	return 400
}
