package errors

type UserAlreadyExistsError struct{}

func (e *UserAlreadyExistsError) Error() string {
	return "E-mail already exists."
}

func (e *UserAlreadyExistsError) StatusCode() int {
	return 409
}
