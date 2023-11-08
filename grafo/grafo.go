package main

import (
	"fmt"
	"revisao-estudos-go/doubleendedqueue"
)

func pesquisa(nome string, grafo map[string][]string) bool {
	filaDePesquisa := doubleendedqueue.NewDoubleEndedQueue()
	filaDePesquisa.PushBack(nome)
	verificadas := make(map[string]bool)

	for !filaDePesquisa.IsEmpty() {
		pessoa := filaDePesquisa.PopFront()

		if _, foiVerificada := verificadas[pessoa]; foiVerificada {
			continue // Pule esta iteração se a pessoa já foi verificada
		}

		if pessoaEhVendedor(pessoa) {
			fmt.Printf("%s é um vendedor de manga!\n", pessoa)
			return true
		}

		for _, contato := range grafo[pessoa] {
			filaDePesquisa.PushBack(contato)
		}

		verificadas[pessoa] = true
	}

	return false
}

func main() {
	grafo := make(map[string][]string)
	grafo = mockGrafo()

	resultado := pesquisa("voce", grafo)
	if !resultado {
		fmt.Println("Nenhum vendedor de manga encontrado.")
	}
}

func mockGrafo() map[string][]string {
	grafo := make(map[string][]string)

	grafo["voce"] = []string{"alice", "bob", "claire"}
	grafo["bob"] = []string{"anuj", "peggy"}
	grafo["alice"] = []string{"peggy"}
	grafo["claire"] = []string{"thom", "jonny"}
	grafo["anuj"] = []string{}
	grafo["peggy"] = []string{}
	grafo["thom"] = []string{}
	grafo["jonny"] = []string{}

	return grafo
}

func pessoaEhVendedor(nome string) bool {
	ultimaLetra := int32(nome[len(nome)-1])
	caractereProcurado := 'm'

	return ultimaLetra == caractereProcurado
}
