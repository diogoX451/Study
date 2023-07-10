package bank

import (
	"fmt"
)

// letras minuscula visivel somente para unio arquvio
// letras  maiuscula visivel para todas importações

type AccountBank struct {
	Titular Titular
	Agencia string
	Numero  int
	saldo   float64
}

func (a *AccountBank) GetSaldo() float64 {
	return a.saldo
}

func (a *AccountBank) Saque(value float64) (float64, error) {
	if value > a.saldo || value < 0 {
		return a.saldo, fmt.Errorf("Valor insuficente para Saque: %f", value)
	}

	a.saldo = a.saldo - value
	return a.saldo, nil
}

func (a *AccountBank) Depositar(value float64) {
	if value < 0 {
		panic("Valor insuficente para Depositar")
	}

	a.saldo = a.saldo + value

}
