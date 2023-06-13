package main

import (
	"fmt"
)

type AccountBank struct {
	agencia string
	numero  int
	saldo   float64
}

func main() {
	var account *AccountBank // Apontar para type
	account = new(AccountBank)
	account.numero = 10
	fmt.Print(account)
}
