package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\nDijkstra em GO\n")
	processados := []string{}

}

func mockGrafoArestasPesos() map[string]map[string]int {
	grafo := make(map[string]map[string]int)

	grafo["inicio"] = make(map[string]int)
	grafo["inicio"]["a"] = 6
	grafo["inicio"]["b"] = 2

	grafo["a"] = make(map[string]int)
	grafo["a"]["fim"] = 1

	grafo["b"] = make(map[string]int)
	grafo["b"]["a"] = 3
	grafo["b"]["fim"] = 5

	grafo["fim"] = make(map[string]int)

	return grafo
}

func mockGrafoCustos() map[string]float64 {
	infinito := math.Inf(1) // Inf(1) representa infinito positivo

	custos := make(map[string]float64)
	custos["a"] = 6
	custos["b"] = 2
	custos["fim"] = infinito
	return custos
}

func mockGrafoPais() map[string]string {

	pais := make(map[string]string)
	pais["a"] = "incio"
	pais["b"] = "incio"
	pais["fim"] = ""
	return pais
}
