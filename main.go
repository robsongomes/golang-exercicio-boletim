package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Nota float32

type Notas [4]Nota

type Aluno struct {
	nome  string
	notas Notas
}

func (a Aluno) String() string {
	line := a.nome + ","
	for _, nota := range a.notas {
		line += fmt.Sprintf("%.2f,", nota)
	}
	return line[:len(line)-1] //remove a última vírgula
}

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

	//carregar alunos do arquivo
	boletim, err := carregarBoletim()

	if err != nil {
		panic(err)
	}

	for {
		opcao := mostrarMenu()

		switch opcao {
		case 1:
			{
				aluno, err := adicionarAluno()
				if err != nil {
					panic(err)
				}
				boletim = append(boletim, aluno)
			}
		case 2:
			imprimirBoletim(boletim)
		default:
			os.Exit(0)
		}

		fmt.Print("\n\n")
	}
}

func converteLinhaParaAluno(linha string) (Aluno, error) {
	campos := strings.Split(linha, ",")
	nome := campos[0]
	notas := Notas{}
	for i := 1; i < len(campos); i++ {
		n, err := strconv.ParseFloat(campos[i], 32)
		if err != nil {
			return Aluno{}, fmt.Errorf("Erro ao converter nota: %w", err)
		}
		notas[i-1] = Nota(n)
	}
	return Aluno{nome, notas}, nil
}

func carregarBoletim() ([]Aluno, error) {
	f, err := os.Open("alunos.txt")
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir arquivo de alunos: %w", err)
	}
	var alunos []Aluno

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		linha := scanner.Text()
		aluno, err := converteLinhaParaAluno(linha)
		if err != nil {
			return nil, fmt.Errorf("Erro ao converter linha: %w", err)
		}
		alunos = append(alunos, aluno)
	}
	return alunos, nil
}

func salvarAlunoNoArquivo(aluno Aluno) error {
	f, err := os.Create("alunos.txt")
	defer f.Close()

	if err != nil {
		return fmt.Errorf("Erro ao criar arquivo: %w", err)
	}
	_, err = fmt.Fprintln(f, aluno) //usa o método String
	if err != nil {
		return fmt.Errorf("Erro ao salvar aluno no arquivo: %w", err)
	}
	return nil
}

/*
Cadastra um novo aluno e suas notas
*/
func adicionarAluno() (Aluno, error) {
	var nome string
	var notas Notas
	fmt.Println("Digite o nome do aluno:")
	fmt.Scanln(&nome)
	fmt.Println("Digite as notas do aluno:")
	for i := 0; i < len(notas); i++ {
		fmt.Printf("Nota %d: ", i+1)
		fmt.Scanln(&notas[i])
	}
	aluno := Aluno{nome, notas}
	err := salvarAlunoNoArquivo(aluno)
	if err != nil {
		return Aluno{}, fmt.Errorf("Erro ao salvar aluno no arquivo: %w", err)
	}
	return aluno, nil
}

/*
Mostra o menu de opções para o usuário
*/
func mostrarMenu() int {
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
func imprimirBoletim(alunos []Aluno) {
	imprimirHeader()
	for _, aluno := range alunos {
		fmt.Printf("%-15s % 7.2f % 7.2f % 7.2f % 7.2f % 7.2f\n",
			aluno.nome,
			aluno.notas[0],
			aluno.notas[1],
			aluno.notas[2],
			aluno.notas[3],
			aluno.notas.media(), //pode colocar em Aluno
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
