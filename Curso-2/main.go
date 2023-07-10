package main

import (
	"fmt"

	"github.com/bank"
)

type Boleto interface {
	Saque(float64) (float64, error)
}

func PaymentBoleto(account Boleto, value float64) {
	account.Saque(value)
}

func main() {
	var account *bank.AccountBank // Apontar para type
	account = new(bank.AccountBank)
	account.Titular.Name = "Diogo Almeida"
	account.Numero = 10
	fmt.Print(account)
	result, err := account.Saque(-10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	//depositar

	account.Depositar(10)
	contaDiogo := bank.AccountBank{}
	contaDiogo.Depositar(1000)
	PaymentBoleto(&contaDiogo, 50)
	fmt.Print(contaDiogo.GetSaldo())
}
