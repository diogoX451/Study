package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Level Package Scopes
/*
	- Não se pode declarar variaveis do tipo abreviado na forma global
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

	nameAndLastname, age := nomeCompleto() // Atribuição multipla

	teste1, _ = nomeCompleto() // Atribuição multipla ignorando o segundo valor

	fmt.Println(nameAndLastname, age)

	var comando int

	fmt.Scanf("%d", &comando) // & referencia da variavel

	level = comando

	fmt.Println(nome, idade)

	fmt.Println(completo)

	Array()
	loop()
	// runTeste()
	teste, err := hasEmbargo()
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println(teste)
	// for { // loop infinito
	// 	switch comando {
	// 	case 1:
	// 		os.Exit(0)
	// 	case 2:
	// 		fmt.Println("2")
	// 	default:
	// 		fmt.Println("Default")
	// 	}

	// }

}

func nomeCompleto() (string, int) {
	completo = nome + lastaname
	idade := 17
	return completo, idade
}

func Array() {
	multiArray := []string{"Diogo", "Huakson"}
	fmt.Println(len(multiArray))
	multiArray = append(multiArray, "Almeida")
	fmt.Println(len(multiArray))
	fmt.Println(cap(multiArray))
	return
}

func loop() {
	teste := []string{"Diogo", "Huakson", "Almeida"}
	for i, teteste := range teste {
		fmt.Println(i, teteste)
	}
}

func runTeste() {
	for i := 0; i < 100000; i++ {
		res, erro := http.Get("http://localhost:8000/")

		if erro != nil {
			fmt.Println(erro)
		}
		if res.StatusCode == 200 {
			fmt.Println(i)
		}
	}
}

func hasEmbargo() (bool, error) {
	arquivo, err := os.Open("../../embargo.csv")
	if err != nil {
		fmt.Println(err)
	}

	doc := bufio.NewScanner(arquivo)

	for doc.Scan() {
		if strings.Contains(doc.Text(), "Adair Borges Pereira") {
			return true, nil
		}
	}

	return false, nil
}
