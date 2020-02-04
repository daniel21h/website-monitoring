package main

import (
		"fmt"
		 "os"
		 "time"//Pacote referente ao tempo
		 "bufio"//Para fazer leitura linha a linha do sites.txt
		 "strings"
		 "net/http"//Pacote para fazer uma requisição web
		 "io"
		 "strconv"//Pacote que converte vários tipos para string
		 "io/ioutil"
)

const monitoramentos = 3
const delay = 5

func main(){
	exibeIntroducao()

	//Fazendo um Looping
	//for sem passagem fica rodando indefinidamente até você cancelar ele
	for {
		exibeMenu()

		comando := leComando()


		//Controle de fluxo com switch
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não cocheço este comando")
			os.Exit(-1)
		}
	}
}


func exibeIntroducao() {
	nome := "Daniel"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
		//Mostrando um menu de opções para o meu usuário
		fmt.Println("1- Iniciar Monitoramento")
		fmt.Println("2- Exibir Logs")
		fmt.Println("0- Sair do Programa")
}

func leComando() int {
	//Capturando o que o nosso usuário irá escrever
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}


func iniciarMonitoramento() {
	fmt.Println("Monitorando....")

	//Slice's em Go
	/*sites := []string {"http://random-status-code.herokuapp.com/",
	 "https://www.alura.com.br", "https://www.caelum.com.br"}*/

	//Lendo um arquivo de texto e fazendo o slice
	sites := leSitesDoArquivo()

	//MONITORAND MÚLTIPLAS VEZES
	for i:= 0; i < monitoramentos ; i++ {
		
		//for i:= 0 ; i < len(sites); i++ {}
		//range é para obter a posição em for e quem está nela
	   for i, site := range sites {
		   fmt.Println("Testando site", i, ":", site)
		   testaSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string){
	resp, err := http.Get(site)
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "esta com problemas. Status code:",
		resp.StatusCode)
		registraLog(site, false)
	}
}

//Trabalhando com arquivos de texto em Go
func leSitesDoArquivo() []string {

	var sites []string 

	//Abrindo e lendo o arquivo
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//O NewReader vai retornar um leitor
	leitor := bufio.NewReader(arquivo)
	for {
		//Irá ler apenas uma linha, usamos for para ler múltiplas linhas
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF{
			break
		}

	}
	//Fechando os arquivos
	arquivo.Close()
	return sites
}

//Abrindo e criando um arquivo(ficou online/offline eu salvo)
func registraLog(site string, status bool) {


	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil{
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	//Convertento o status(bool) to string

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}