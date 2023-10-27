package main

import "fmt"

func main() {
	fmt.Println("\nTabelas Hash/Map em GO\n")
	demonstracaoMap()
}

func demonstracaoMap() {
	caderno := make(map[string]float64)

	caderno["maçã"] = 0.64
	caderno["leite"] = 1.49
	caderno["abacate"] = 1.49

	fmt.Println(caderno)
}
