package main

import "fmt"

func main() {
	// var nome string = "Roberto"
	// var versao float32 = 1.18
	// var idade int = 29

	nome := "Roberto"
	versao := 1.18

	fmt.Println("Hellooooo, ", nome)
	fmt.Println("Este programa esta na versao, ", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	// "%d" representa um inteiro
	// o Scanf recebe um modificado "%d" e atribui à variavel comando
	// o & antes do &comando representa o endereço da variavel, um ponteiro para a variavel
	//fmt.Scanf("%d", &comando)
	// Scan recebe um int
	fmt.Scan(&comando)
	fmt.Println("O endereço da variavel comando é: ", &comando)
	fmt.Println("O comando escolhido foi ", comando)

}
