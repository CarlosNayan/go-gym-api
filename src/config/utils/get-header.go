package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetHeader é uma função genérica que tenta pegar um cabeçalho e convertê-lo para int.
// Caso o cabeçalho não exista ou a conversão falhe, o valor padrão (defaultValue) é retornado.
func GetHeader[T int | string](ctx *fiber.Ctx, headerName string, defaultValue T) T {
	headerValue := ctx.Get(headerName)

	if headerValue == "" {
		return defaultValue
	}

	var result T

	switch any(defaultValue).(type) {
	case int:
		intVal, err := strconv.Atoi(headerValue)
		if err != nil {
			return defaultValue
		}
		result = any(intVal).(T)
	case string:
		result = any(headerValue).(T)
	default:
		return defaultValue
	}

	

	return result
}
