package main

import (
	"fmt"
	"net/http"
)

type Aluno struct {
	name     string
	age      int
	lastname string
}

var constructAluno *Aluno

func main() {

	res, error := http.Get("https://google.com")
	if res.Body != nil {
		defer fmt.Println(res.Cookies()) // defer delay of execution
	}
	fmt.Println(res.StatusCode, error)

	constructAluno = new(Aluno)
	constructAluno.name = "Diogo"
	constructAluno.age = 17
	constructAluno.lastname = "Almeida"

	fmt.Println(constructAluno)
}
