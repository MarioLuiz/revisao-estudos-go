package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Mario"
	idade := 30
	versao := 1.0
	fmt.Println("\n")
	fmt.Println("Ola Sr. ", nome, " sua idade é ", idade)
	fmt.Println("Programa na versao ", versao)
	fmt.Println("\n")

	fmt.Println("O tipo da variavel versao e ", reflect.TypeOf(versao))
	fmt.Println("O tipo da variavel idade e ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao e ", reflect.TypeOf(nome))
}
