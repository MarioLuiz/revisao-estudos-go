package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("\nRecursao em GO\n")
	pilha, quantidadeCaixas := criaUmaPilhaParaBusca()
	fmt.Println("Quantidade de caixas: ", quantidadeCaixas)
	fmt.Println("Pilha de Caixas: \n", pilha)

	procuraPelaChaveComLoops(pilha)

	start := time.Now()
	procuraPelaChaveComRecursao(pilha, 0)
	elapsed := time.Since(start)
	fmt.Println("Procurar pela Chava utilizando recirsao demorou: ", elapsed)

}

// Problema: Existe uma chave que abre uma mala misteriosa. A sua avó diz que a chave está em uma caixa. Esta caixa contém mais caixas com mais caixas dentro delas.
// Utilizando Loops
func procuraPelaChaveComLoops(pilha [][]string) {
	start := time.Now()
	fmt.Println("\n")
	achouChave := false
	contador := 0
	for i := 0; achouChave == false; i++ {
		caixas := pilha[i]
		for i := 0; i < len(caixas); i++ {
			contador++
			if caixas[i] == "Vazia" {
				caixa := []string{caixas[i]}
				pilha = append(pilha, caixa)
			} else {
				fmt.Println("Achou a chave na volta: ", contador)
				achouChave = true
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Como ficou a Pilha: \n", pilha)
	fmt.Println("Procurar pela Chava utilizando loops demorou: ", elapsed)
}

func procuraPelaChaveComRecursao(pilha [][]string, contador int) {
	fmt.Println("\n")
	if len(pilha) == 0 {
		fmt.Println("A chave não foi encontrada na pilha.")
		return
	}
	caixas := pilha[0]
	for i := 0; i < len(caixas); i++ {
		contador++
		if caixas[i] == "Chave" {
			fmt.Println("Achou a chave na volta:", contador)
			fmt.Println("Como ficou a Pilha: \n", pilha)
			return
		}
	}
	procuraPelaChaveComRecursao(pilha[1:], contador)
}

func criaUmaPilhaParaBusca() ([][]string, int) {
	quantidadeDeCaixas := rand.Intn(15-10) + 10
	quantidadeCaixasDentroDaCaixa := rand.Intn(5-1) + 1
	randomizando := quantidadeCaixasDentroDaCaixa * quantidadeDeCaixas
	caixaQueTemChave := rand.Intn(randomizando)
	fmt.Println("Caixa que tem a chave numero: ", caixaQueTemChave)
	contadorDeCaixas := 0
	pilhaDeCaixas := [][]string{}
	for i := 0; i < quantidadeDeCaixas; i++ {
		pilhaDeCaixa := []string{}
		for i := 0; i < quantidadeCaixasDentroDaCaixa; i++ {
			contadorDeCaixas++
			if contadorDeCaixas == caixaQueTemChave {
				pilhaDeCaixa = append(pilhaDeCaixa, "Chave")
				caixaQueTemChave--
			} else {
				pilhaDeCaixa = append(pilhaDeCaixa, "Vazia")
			}
		}
		pilhaDeCaixas = append(pilhaDeCaixas, pilhaDeCaixa)
	}
	caixaQueTemChave++
	return pilhaDeCaixas, contadorDeCaixas
}
