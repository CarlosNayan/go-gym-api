package utils

import (
	"fmt"
	"runtime"
)

func WrapError(err error) error {
	_, file, line, _ := runtime.Caller(1) // Captura o chamador
	return fmt.Errorf("error at %s:%d: %w", file, line, err)
}
