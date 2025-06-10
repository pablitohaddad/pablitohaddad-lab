package main

import (
	"log"
	"os"
	"using-gemini-api/email"
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

	respostaGemini, err := gemini.AnalyzeText(apiKey, mockDescricoesDeVaga())
	if err != nil {
		log.Fatalf("Erro na chamada da API do Gemini")
	}
	
	email.SendEmail(respostaGemini)

}

func mockDescricoesDeVaga() string{
	return  `
		Vaga: Estágio de Desenvolvimento Backend
		Empresa: Tech Solutions Inc.
		Local: Remoto (Brasil)
		Publicado em: 07/06/2025
		Link para se inscrever: https://techsolutions.com/vaga/123
		
		Descrição:
		Estamos em busca de um Estagiário de Desenvolvimento Backend para se juntar ao nosso time de inovação.
		O candidato ideal está cursando Ciência da Computação ou áreas correlatas e tem paixão por tecnologia.
		Você irá trabalhar com microsserviços em Go, bancos de dados PostgreSQL e Docker.

		Vaga: Estágio em Engenharia de Software (Java & Cloud)
   	 	Empresa: Cloud Innovators Ltda.
    	Local: São Paulo - SP (Híbrido)
    	Publicado em: 09/06/2025
    	Link para se inscrever: https://cloudinnovators.com/careers/java-intern
    
    	Descrição:
    	A Cloud Innovators busca um(a) Estagiário(a) de Engenharia de Software para integrar nosso time de microsserviços.
    	Você irá nos ajudar a construir e escalar APIs RESTful de alta performance usando Java e o framework Spring Boot.
    	Nossa stack utiliza MongoDB como banco de dados NoSQL principal e Redis para gerenciamento de cache de alta velocidade. 
    	Toda a nossa infraestrutura é baseada em AWS, e você terá a oportunidade de aprender sobre serviços essenciais da nuvem.

		Vaga: Desenvolvedor(a) de Software Sênior - Foco em Python
    	Empresa: Fintech Global
    	Local: Belo Horizonte - MG (Presencial)
    	Publicado em: 09/06/2025
    	Link para se inscrever: https://fintechglobal.com/jobs/senior-python-dev

    	Descrição:
    	Buscamos um(a) Desenvolvedor(a) Sênior com sólida experiência para liderar tecnicamente nossa equipe de pagamentos.
    	É necessário ter pelo menos 5 anos de experiência com desenvolvimento de software e arquitetura de sistemas distribuídos.
    	O candidato será responsável por projetar novas soluções, definir padrões de código, mentorar desenvolvedores mais jovens e garantir a escalabilidade de nossas aplicações.
    	Requisitos: Profundo conhecimento em Python, Django, arquitetura de microsserviços e experiência com Kafka.
	`
}