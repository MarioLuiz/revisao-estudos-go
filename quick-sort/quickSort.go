package main

import "fmt"

func main() {
	fmt.Println("\nQuickSort em GO\n")
	array := []int{10, 5, 2, 3}
	resultado := quicksort(array)
	fmt.Println(resultado)
}

func quicksort(arrayDesordenado []int) []int {
	if len(arrayDesordenado) < 2 {
		return arrayDesordenado
	} else {
		pivo := arrayDesordenado[0]
		menores := []int{}
		maiores := []int{}
		for i := 1; i < len(arrayDesordenado); i++ {
			if arrayDesordenado[i] < pivo {
				menores = append(menores, arrayDesordenado[i])
			} else {
				maiores = append(maiores, arrayDesordenado[i])
			}
		}
		menores = quicksort(menores)
		maiores = quicksort(maiores)
		return append(append(menores, pivo), maiores...)
	}
}
