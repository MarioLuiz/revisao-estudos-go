package main

import "fmt"

func main() {
	fmt.Println("Tabelas Hash/Map Votacao em Go")
	votacao()
}

func votacao() {
	listaVotacao := mockListaVotacao()
	verificaEleitorPodeVotar(listaVotacao, "Mario")
	verificaEleitorPodeVotar(listaVotacao, "Fabio")
	verificaEleitorPodeVotar(listaVotacao, "Carla")
	verificaEleitorPodeVotar(listaVotacao, "Beatriz")
	verificaEleitorPodeVotar(listaVotacao, "Luiz")
	verificaEleitorPodeVotar(listaVotacao, "Roberto")
	verificaEleitorPodeVotar(listaVotacao, "Maria")
	verificaEleitorPodeVotar(listaVotacao, "Katia")
	verificaEleitorPodeVotar(listaVotacao, "Revihery")
	verificaEleitorPodeVotar(listaVotacao, "Paulo")
}

func verificaEleitorPodeVotar(listaVotacao map[string]string, nomeEleitor string) {
	cpf, ok := listaVotacao[nomeEleitor]
	if ok {
		fmt.Println("Mande embora! Eleitor: " + nomeEleitor + " CPF: " + cpf)
	} else {
		fmt.Println("Deixe votar! Eleitor: " + nomeEleitor)
	}
}

func mockListaVotacao() map[string]string {
	listaVotacao := make(map[string]string)

	listaVotacao["Mario"] = "04319232199"
	listaVotacao["Luiz"] = "33319232199"
	listaVotacao["Carla"] = "96548768132"
	listaVotacao["Maria"] = "04319232199"
	listaVotacao["Revihery"] = "02719232159"

	return listaVotacao
}
