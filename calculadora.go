package calculadora

import "fmt"

// Soma retorna a soma de dois números
func Soma(a, b float64) float64 {
	return a + b
}

// Subtrai retorna a subtração de dois números
func Subtrai(a, b float64) float64 {
	return a - b
}

// Multiplica retorna a multiplicação de dois números
func Multiplica(a, b float64) float64 {
	return a * b
}

// Divide retorna a divisão de dois números
// Se o divisor for 0, retorna um erro
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("não é possível dividir por zero")
	}
	return a / b, nil
}
