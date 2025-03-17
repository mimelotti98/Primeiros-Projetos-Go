package calculadora

import (
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := Soma(2, 3)
	esperado := 5.0
	if resultado != esperado {
		t.Errorf("Soma(2, 3) = %v; esperado %v", resultado, esperado)
	}
}

func TestSubtrai(t *testing.T) {
	resultado := Subtrai(5, 3)
	esperado := 2.0
	if resultado != esperado {
		t.Errorf("Subtrai(5, 3) = %v; esperado %v", resultado, esperado)
	}
}

func TestMultiplica(t *testing.T) {
	resultado := Multiplica(2, 3)
	esperado := 6.0
	if resultado != esperado {
		t.Errorf("Multiplica(2, 3) = %v; esperado %v", resultado, esperado)
	}
}

func TestDivide(t *testing.T) {
	resultado, err := Divide(6, 3)
	esperado := 2.0
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado != esperado {
		t.Errorf("Divide(6, 3) = %v; esperado %v", resultado, esperado)
	}
}

func TestDividePorZero(t *testing.T) {
	_, err := Divide(6, 0)
	if err == nil {
		t.Error("Esperado erro ao dividir por zero, mas não ocorreu.")
	}
}
