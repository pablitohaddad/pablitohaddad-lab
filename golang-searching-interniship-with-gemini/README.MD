# Caça-Vagas com IA usando Go e Gemini

[![Go](https://img.shields.io/badge/Go-1.22-blue.svg)](https://go.dev/)
[![Gemini](https://img.shields.io/badge/Gemini-1.5%20Flash-purple.svg)](https://ai.google.dev/)
[![Chromedp](https://img.shields.io/badge/Chromedp-v0.9-brightgreen.svg)](https://github.com/chromedp/chromedp)

Um projeto que automatiza a busca por vagas de estágio em tecnologia, utilizando web scraping para coletar os dados e a API do Gemini para analisar e extrair informações relevantes de forma inteligente. Este README cobre a implementação até a análise com IA.

## 🧠 Conceitos Principais

Este projeto combina duas etapas fundamentais que trabalham em sequência para transformar dados não estruturados da web em informações úteis e organizadas.

1.  **Scraper (O Coletor):** Navega em páginas web de vagas (usando `chromedp` para sites dinâmicos) e extrai o texto bruto das descrições. Ele atua como nossos "olhos" na internet. (PS: Ainda não implementado!!)

2.  **Analyzer (O Cérebro com IA):** Recebe o texto bruto do Coletor e o envia para a API do Gemini com um prompt específico, pedindo para filtrar, analisar e retornar um objeto JSON estruturado com as informações chave da vaga.

## 📦 Como Rodar o Projeto (Até a Integração com Gemini)

Siga os passos abaixo para configurar e executar a aplicação no seu ambiente local.

### Pré-requisitos
* **Go (versão 1.21+):** [Instruções de Instalação](https://go.dev/doc/install)
* **API Key do Gemini:** Crie sua chave no [Google AI Studio](https://aistudio.google.com/)

### Passos para a execução

```bash
# 1. Clone o repositório (substitua pela URL do seu repositório)
git clone [https://github.com/seu-usuario/seu-repositorio.git](https://github.com/seu-usuario/seu-repositorio.git)

# 2. Entre na pasta do projeto
cd seu-repositorio

# 3. Instale as dependências
go mod tidy

# 4. Configure sua API Key
# Crie um arquivo chamado .env na raiz do projeto e adicione sua chave.
# Lembre-se de adicionar o arquivo .env ao seu .gitignore!
touch .env
echo 'GEMINI_API_KEY="COLE_SUA_CHAVE_DA_API_AQUI"' > .env

# 5. Execute o programa principal
# Ele irá rodar o scraper com um exemplo e imprimir a análise do Gemini no terminal.
go run main.go
```

### ✅ Resultado Esperado

Após a execução, você verá um JSON formatado no seu terminal, que é a resposta da análise do Gemini sobre a vaga encontrada.

```json
{
  "job_title": "Estágio de Desenvolvimento Backend",
  "company": "Tech Solutions Inc.",
  "role": "Backend",
  "key_technologies": [
    "Go",
    "PostgreSQL",
    "Docker"
  ],
  "location": "Remoto (Brasil)",
  "publication_date": "08/06/2025",
  "job_link": "[https://techsolutions.com/vaga/123](https://techsolutions.com/vaga/123)",
  "is_internship": true
}
```