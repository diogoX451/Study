package main

import (
	"fmt"
	"net/http"
)

func main() {

	res, error := http.Get("https://google.com")
	fmt.Println(res, error)
}
