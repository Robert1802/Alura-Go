package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 2

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo...")
		// } else {
		// 	fmt.Println("Não conheço este comando")
		// }

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {

	// var nome string = "Roberto"
	// var versao float32 = 1.18
	// var idade int = 29

	nome := "Roberto"
	versao := 1.18

	fmt.Println("Hellooooo, ", nome)
	fmt.Println("Este programa esta na versao, ", versao)

}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	// "%d" representa um inteiro
	// o Scanf recebe um modificado "%d" e atribui à variavel comandoLido
	// o & antes do &comandoLido representa o endereço da variavel, um ponteiro para a variavel
	//fmt.Scanf("%d", &comandoLido)
	// Scan recebe um int
	fmt.Scan(&comandoLido)
	//fmt.Println("O endereço da variavel comandoLido é: ", &comandoLido)
	fmt.Println("O comandoLido escolhido foi ", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	// for i := 0; i < len(sites); i++ {
	// ou
	// for i, site := range sites {

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site [", site, "] que está na posição ", i, " do meu Slice")
			testaSite(site)
		}
		// Aguarda 2 segundos antes do proximo monitoramento
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("O erro foi: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado corretamente")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, " foi está com problemas. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func exemploDeSlice() {
	//Isso é um Slice, uma abstração da array do Go com tamanho variavel
	nomes := []string{"Nome 1", "Nome 2", "Nome 3"}
	fmt.Println(nomes)
	fmt.Println("O meu Slice tem:", len(nomes), " itens")
	fmt.Println("A capacidade do Slice é:", cap(nomes), " itens")

	// O slice dobra a sua capacidade antes de estourar com append
	nomes = append(nomes, "Nome 4")
	fmt.Println(nomes)
	fmt.Println("O meu slice tem:", len(nomes), " itens")
	fmt.Println("A capacidade do Slice é:", cap(nomes), " itens")

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("O erro foi: ", err)
	}

	// bufio vai ler o arquivo linha a linha
	leitor := bufio.NewReader(arquivo)

	for {
		// a quebra de linha ocorre no \n que é a quebra de linha
		linha, err := leitor.ReadString('\n')
		// retira o Enter extra ou os espaços sobrando
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
