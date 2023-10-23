package main

import "fmt"

/*
Você quer dividir sua fazenda em porções quadradas iguais, sendo que estas
porções devem ter o maior tamanho possível.
*/
func main() {
	fmt.Println("\n Procurando pedaço terreno com quadrados iguais")
	comprimento := 1680
	largura := 640
	comprimento, largura = divideTerrenoPorcaoQuadradasIguais(comprimento, largura)
	if comprimento != largura {
		fmt.Println("\nNão foi possivel achar a porção de terrenos iguais")
	} else {
		fmt.Println("\nAs porções de terras iguais ficaram com as medidas: %d %d", comprimento, largura)
	}

}
func divideTerrenoPorcaoQuadradasIguais(comprimento int, largura int) (int, int) {
	if eUmCasoBase(comprimento, largura) {
		return comprimento, largura
	}
	largura, comprimento = MenorLadoAndMaiorLado(comprimento, largura)
	porcaoDeTerraNaoDividida := 0

	if (comprimento * 2) < largura {
		porcaoDeTerraNaoDividida = comprimento - (largura * 2)
		comprimento = largura
		largura = porcaoDeTerraNaoDividida
		return divideTerrenoPorcaoQuadradasIguais(comprimento, largura)
	} else if comprimento > largura {
		porcaoDeTerraNaoDividida = comprimento - largura
		comprimento = largura
		largura = porcaoDeTerraNaoDividida
		return divideTerrenoPorcaoQuadradasIguais(comprimento, largura)
	}
	return divideTerrenoPorcaoQuadradasIguais(comprimento, largura)
}

func MenorLadoAndMaiorLado(comprimento int, largura int) (int, int) {
	manorLado := comprimento
	maiorLado := largura
	if comprimento < largura {
		manorLado = comprimento
		maiorLado = largura
	} else {
		manorLado = largura
		maiorLado = comprimento
	}
	return manorLado, maiorLado
}

func eUmCasoBase(comprimento int, largura int) bool {
	if comprimento == largura {
		return true
	} else {
		return false
	}
}
