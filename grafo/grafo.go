package main

import (
	"fmt"
	"revisao-estudos-go/doubleendedqueue"
)

func main() {
	fmt.Println("Grafo em Go")
	deque := doubleendedqueue.NewDoubleEndedQueue()

	deque.PushBack("apple")
	deque.PushFront("banana")
	deque.PushBack("cherry")

	fmt.Println("Front:", deque.PopFront())
	fmt.Println("Back:", deque.PopBack())
	fmt.Println("Front:", deque.PopFront())
	fmt.Println("Front:", deque.PopFront())
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
