package main

import (
	"fmt"
	"reflect"
)

// Level Package Scopes
/*
	- Não se pode declarar variaveis do tipo abriviado na forma global
		so funciona em code blocks
*/

var level int = 10
var completo string
var nome string = "Diogo"
var lastaname string = "Almeida"

func main() {

	var idade int = 15
	var preco float32 = 1.1
	var veriffyType = reflect.TypeOf(preco) // Verificar o type
	teste1 := "Guilherme"                   // Operador curto ja define automaticamente o type
	fmt.Println(veriffyType, nome, idade, preco, teste1)

	nome = "Huakson"

	if idade != 15 {
		parabens := "Parabens"
		fmt.Println(parabens)
	}

	var comando int

	fmt.Scanf("%d", &idade) // & referencia da variavel

	level = comando

	fmt.Println(nome, idade)

	fmt.Println(completo)

}

func nomeCompleto() {
	completo = nome + lastaname
	fmt.Println(completo)
	return
}
