package main

import (
	"fmt"
	"os"
	"strconv"
)

type Nota float32

type Notas [4]Nota

type Boletim map[string]Notas

/*
Calcula a média de 4 notas
*/
func (n Notas) media() Nota {
	var media Nota
	for _, nota := range n {
		media += nota
	}
	return media / Nota(len(n))
}


func main() {
	
	var boletim = Boletim {
		"João Gomes": {8, 7, 9, 10},
		"André": {8, 5, 7.55, 10},
		"Marta": {10,10,10,10},
	}

	for {
		opcao := mostrarMenu()

		switch opcao {
			case 1: {
				nome, notas := adicionarAluno()
				boletim[nome] = notas
			}
			case 2: imprimirBoletim(boletim)
			default: os.Exit(0)
		}

		fmt.Print("\n\n")
	}
}

/*
Cadastra um novo aluno e suas notas
*/
func adicionarAluno() (string, Notas) {
	var nome string
	var notas Notas
	fmt.Println("Digite o nome do aluno:")
	fmt.Scanln(&nome)
	fmt.Println("Digite as notas do aluno:")
	for i := 0; i < len(notas); i++ {
		fmt.Printf("Nota %d: ", i+1)
		fmt.Scanln(&notas[i])
	}
	return nome, notas
}

/*
Mostra o menu de opções para o usuário
*/
func mostrarMenu() int{
	var opcao string	
	fmt.Println("1 - Adicionar aluno")
	fmt.Println("2 - Imprimir Boletim")
	fmt.Println("3 - Encerrar")
	fmt.Scanln(&opcao)
	op, _ := strconv.Atoi(opcao)
	return op
}

/*
Imprime o boletim na tela em forma de tabela
*/
func imprimirBoletim(alunos Boletim) {	
	imprimirHeader()
	for nome, notas := range alunos {
		fmt.Printf("%-15s % 7.2f % 7.2f % 7.2f % 7.2f % 7.2f\n", 
			nome, 
			notas[0], 
			notas[1], 
			notas[2], 
			notas[3],
			// calcularMedia(notas),
			notas.media(),
		)
	}
}

/*
Imprime o cabeçalho do boletim
*/
func imprimirHeader() {
	fmt.Printf("%-15s %7s %7s %7s %7s %7s\n", "NOME", "AV1", "AV2", "AV3", "AV4", "MÉDIA")
	fmt.Println("--------------------------------------------------------")
}

// func calcularMedia(notas Notas) Nota {
// 	var media Nota
// 	for _, nota := range notas {
// 		media += nota
// 	}
// 	return media / Nota(len(notas))
// }