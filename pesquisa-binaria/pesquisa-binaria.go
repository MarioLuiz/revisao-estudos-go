package main

import "fmt"

func main() {
	itemProcurado := 16
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	achouItemProcurado, contador := pesquisaBinaria(itemProcurado, lista)
	if achouItemProcurado {
		fmt.Println("\n Pesquisa achou o item procurado na volta ", contador)
	} else {
		fmt.Println("\n Pesquisa não achou o item procurado ", itemProcurado)
	}
}

func pesquisaBinaria(itemProcurado int, lista []int) (bool, int) {
	fmt.Println("\n Algoritmo Pesquisa Binaria em GO")
	baixo := 0
	alto := len(lista) - 1
	contador := 0
	chute := 0
	for baixo <= alto {
		contador++
		var meio int = (baixo + alto) / 2
		chute := lista[meio]
		if chute == itemProcurado {
			return true, contador
		}
		if chute > itemProcurado {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return false, chute
}
