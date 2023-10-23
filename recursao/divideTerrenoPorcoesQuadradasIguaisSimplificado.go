package main

import "fmt"

/*
Você quer dividir sua fazenda em porções quadradas iguais, sendo que estas
porções devem ter o maior tamanho possível.
*/
func main() {
	fmt.Println("\nProcurando pedaço de terreno com quadrados iguais")
	comprimento := 1680
	largura := 640
	comprimento, largura = divideTerrenoPorcoesQuadradasIguais(comprimento, largura)
	if comprimento == largura {
		fmt.Printf("\nAs porções de terras iguais ficaram com as medidas: %d %d\n", comprimento, largura)
	} else {
		fmt.Println("\nNão foi possível encontrar uma divisão com porções de terras iguais")
	}
}

func divideTerrenoPorcoesQuadradasIguais(comprimento int, largura int) (int, int) {
	if casoBase(comprimento, largura) {
		return comprimento, largura
	}

	if comprimento > largura {
		return divideTerrenoPorcoesQuadradasIguais(comprimento-largura, largura)
	} else {
		return divideTerrenoPorcoesQuadradasIguais(comprimento, largura-comprimento)
	}
}

func casoBase(comprimento int, largura int) bool {
	return comprimento == largura
}
