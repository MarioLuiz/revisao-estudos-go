package main

import "fmt"

func main() {
	listaDesordenada := []int{22, 64, 1, 35, 15, 21, 66, 37, 87, 182, 222, 133, 122, 177, 358, 154, 321, 187, 98, 123}
	listaOrdenada := []int{}
	fmt.Println("\n Algoritmo ordenacao por selecao em GO")
	for len(listaDesordenada) > 0 {
		IndiceComMenorValorEncontrado := buscaIndiceComMenorValor(listaDesordenada)
		// Adicionando menor valor em um novo Slice
		listaOrdenada = append(listaOrdenada, listaDesordenada[IndiceComMenorValorEncontrado])
		// Removendo indice do Slice desordenado
		listaDesordenada = append(listaDesordenada[:IndiceComMenorValorEncontrado], listaDesordenada[IndiceComMenorValorEncontrado+1:]...)
	}
	fmt.Println("Lista Desordenada: ", listaDesordenada)
	fmt.Print("\n")
	fmt.Println("ListaOrdenada", listaOrdenada)
}

func buscaIndiceComMenorValor(lista []int) int {
	menor := lista[0]
	menorIndice := 0
	for i := 0; i < len(lista); i++ {
		if lista[i] < menor {
			menor = lista[i]
			menorIndice = i
		}
	}
	return menorIndice
}
