package main

import "fmt"

func main() {
	fmt.Println("Grafo em Go")
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
	return nome[-1] == 'm'
}
