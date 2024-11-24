package errors

import "fmt"

// UserAlreadyExistsError é um erro para quando o e-mail já existe
type UserAlreadyExistsError struct{}

func (e *UserAlreadyExistsError) Error() string {
	return "E-mail already exists."
}

// InvalidCredentialsError é um erro para credenciais inválidas
type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "Invalid credentials error."
}

// ResourceNotFoundError é um erro para recurso não encontrado
type ResourceNotFoundError struct{}

func (e *ResourceNotFoundError) Error() string {
	return "Resource not found."
}

// MaxDistanceError é um erro para distância máxima atingida
type MaxDistanceError struct{}

func (e *MaxDistanceError) Error() string {
	return "Too far to check in on selected gym!"
}

// MaxNumberOfCheckInsError é um erro para limite de check-ins atingido
type MaxNumberOfCheckInsError struct{}

func (e *MaxNumberOfCheckInsError) Error() string {
	return "User already checked in today!"
}

// LateCheckinValidationError é um erro para validação tardia de check-in
type LateCheckinValidationError struct{}

func (e *LateCheckinValidationError) Error() string {
	return "The check-in can only be validated until 20 minutes of its creation!"
}

// CustomError com mais informações, se necessário
type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
