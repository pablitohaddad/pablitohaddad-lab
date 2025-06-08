package main

import (
	"fmt"
	"log"
	"os"
	"using-gemini-api/gemini"

	"github.com/joho/godotenv"
)

func main(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }
    apiKey := os.Getenv("GEMINI_API_KEY")
    if apiKey == "" {
        log.Fatal("A variável de ambiente GEMINI_API_KEY não foi definida.")
    }

	descricaoDeVaga := `
		Vaga: Estágio de Desenvolvimento Backend
		Empresa: Tech Solutions Inc.
		Local: Remoto (Brasil)
		Publicado em: 07/06/2025
		Link para se inscrever: https://techsolutions.com/vaga/123
		
		Descrição:
		Estamos em busca de um Estagiário de Desenvolvimento Backend para se juntar ao nosso time de inovação.
		O candidato ideal está cursando Ciência da Computação ou áreas correlatas e tem paixão por tecnologia.
		Você irá trabalhar com microsserviços em Go, bancos de dados PostgreSQL e Docker.
	`
	respostaGemini, err := gemini.AnalyzeText(apiKey, descricaoDeVaga)
	if err != nil {
		log.Fatalf("Erro na chamada da API do Gemini")
	}
	fmt.Println("Resposta do Gemini: ", respostaGemini)

}