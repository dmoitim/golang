package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const qtdeMonitoramentos = 3
const segundosEspera = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "DevA"
	versao := 1.1

	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"http://random-status-code.herokuapp.com/",
		"https://compras.empro.com.br/",
		"https://www.empro.com.br/",
		"https://www.riopreto.sp.gov.br/",
		"https://gestao.riopreto.sp.gov.br/"}

	// Testa mais de uma vez
	for i := 0; i < qtdeMonitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}

		// Espera antes da próxima execução
		time.Sleep(segundosEspera * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problema. Status Code:", resp.StatusCode)
	}
}
