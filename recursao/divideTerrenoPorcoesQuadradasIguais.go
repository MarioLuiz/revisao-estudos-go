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

}
func divideTerrenoPorcaoQuadradasIguais(comprimento int, largura int) (int, int) {
	menorLado, maiorLado := MenorLadoAndMaiorLado(comprimento, largura)
	porcaoDeTerraNaoDividida := 0

	if eUmQuadrado(maiorLado, menorLado) {
		return maiorLado, menorLado
	} else if (menorLado * 2) < maiorLado {
		porcaoDeTerraNaoDividida = (menorLado * 2) - maiorLado
		maiorLado = menorLado
		menorLado = porcaoDeTerraNaoDividida
		divideTerrenoPorcaoQuadradasIguais(maiorLado, menorLado)
	}
	return 0, 0
}

func MenorLadoAndMaiorLado(comprimento int, largura int) (int, int) {
	manorLado := 0
	maiorLado := 0
	if comprimento < largura {
		manorLado = comprimento
		maiorLado = largura
	} else {
		manorLado = largura
		maiorLado = comprimento
	}
	return manorLado, maiorLado
}

func eUmQuadrado(comprimento int, largura int) bool {
	if (comprimento % largura) == 0 {
		return true
	} else {
		return false
	}
}
