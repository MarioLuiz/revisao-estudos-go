package main

import (
	"fmt"
	"math/big"
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
	fmt.Println("Procurar pela Chava utilizando recursao demorou: ", elapsed)

	numeroFatorialCalculado := 350
	resultadoFatorial := big.NewInt(0)
	fmt.Println("\nFatorial com recursao e com Loops, numero calculado: ", numeroFatorialCalculado)

	resultadoFatorial = fatorialComLoop(numeroFatorialCalculado)
	fmt.Println("Resultado calculo faltorial com Loops: ", resultadoFatorial)

	fmt.Println("\n")

	init := time.Now()
	resultadoFatorial = fatorialRecursivo(big.NewInt(int64(numeroFatorialCalculado)))
	finit := time.Since(init)
	fmt.Println("Calcular Fatorial recursao demorou: ", finit)
	fmt.Println("Resultado calculo faltorial com Loops: ", resultadoFatorial)

}

// Problema: Existe uma chave que abre uma mala misteriosa. A sua avó diz que a chave está em uma caixa. Esta caixa contém mais caixas com mais caixas dentro delas.
// Utilizando Loops
func procuraPelaChaveComLoops(pilha [][]string) {
	start := time.Now()
	fmt.Println("\n")

	achouChave := false
	contador := 0

	for i := 0; !achouChave && i < len(pilha); i++ {
		caixas := pilha[i]

		for j := 0; j < len(caixas); j++ {
			contador++
			if caixas[j] == "Vazia" {
				caixa := []string{caixas[j]}
				pilha = append(pilha, caixa)
			} else {
				fmt.Println("Achou a chave na volta: ", contador)
				achouChave = true
				break
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Como ficou a Pilha:\n", pilha)
	fmt.Println("Procurar pela Chave utilizando loops demorou: ", elapsed)
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

func fatorialRecursivo(x *big.Int) *big.Int {
	if x.Cmp(big.NewInt(1)) <= 0 {
		return big.NewInt(1)
	}

	temp := new(big.Int)
	temp.Set(x)
	temp.Sub(temp, big.NewInt(1))
	return temp.Mul(x, fatorialRecursivo(temp))
}

func fatorialComLoop(x int) *big.Int {
	start := time.Now()
	resultado := big.NewInt(1)
	for i := 1; i <= x; i++ {
		resultado.Mul(resultado, big.NewInt(int64(i)))
	}
	elapsed := time.Since(start)
	fmt.Println("Calcular Fatorial loops demorou: ", elapsed)
	return resultado
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
