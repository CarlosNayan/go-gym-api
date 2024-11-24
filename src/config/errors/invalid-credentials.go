package errors

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "Invalid credentials."
}

func (e *InvalidCredentialsError) StatusCode() int {
	return 401
}
