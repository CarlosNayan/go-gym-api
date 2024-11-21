package utils

import (
	"api-gym-on-go/src/config/utils"
	"testing"
)

// TestSum testa a função Sum
func TestSum(t *testing.T) {
	t.Run("Testa soma de dois números positivos", func(t *testing.T) {
		result := utils.Sum(2, 3)
		if result != 5 {
			t.Errorf("Esperado 5, mas obteve %d", result)
		}
	})

	t.Run("Testa soma de números negativos", func(t *testing.T) {
		result := utils.Sum(-2, -3)
		if result != -5 {
			t.Errorf("Esperado -5, mas obteve %d", result)
		}
	})
}
