package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AnalyzeText(apiKey, textToAnalyze string) (string, error){
	// Conexão com o Gemini
	ctx := context.Background()
    client, err := genai.NewClient(ctx ,option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatal(err)
    }
	defer client.Close()

	// Qual modelo iremos utilizar e instruções para ele
	model := client.GenerativeModel("gemini-2.0-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("Você é um assistente de RH para tecnologia, focado em vagas de estágio no Brasil. Sua tarefa é analisar o texto de uma vaga e extrair as informações solicitadas no formato JSON. Responda apenas com o JSON."),
		},
	}

  	prompt := fmt.Sprintf(`
		--- DESCRIÇÃO DA VAGA ---
		%s
		--- FIM DA DESCRIÇÃO ---

		Responda APENAS com um objeto JSON válido, com a seguinte estrutura:
		{
		  "titulo_vaga": "O nome da vaga",
		  "empresa": "O nome da empresa",
		  "cargo": "A função principal (ex: Backend, Frontend, Full Stack)",
		  "tecnologias_principais": ["Uma", "lista", "de", "tecnologias"],
		  "local": "O local da vaga (ex: Remoto, São Paulo - SP)",
		  "data_publicacao": "A data em que a vaga foi postada",
		  "url_vafa": "O link para se inscrever na vaga",
		  "is_estagio": true
		}
	`, textToAnalyze)

	// Pegando a resposta do modelo
    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return "", fmt.Errorf("erro ao gerar conteúdo: %v", err)
    }
    
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("resposta inesperada ou vazia do Gemini")
	}

	// Pega a primeira resposta do modelo, que ele gerou. 
	// A IA geralmente possuim várias respostas candidatas, nós pegamos a primeira
	part := resp.Candidates[0].Content.Parts[0]
	
	// Garantindo que o modelo me retornou uma String, de fato
	if text, ok := part.(genai.Text); ok {
		return string(text), nil
	}

	return "", fmt.Errorf("resposta não contém texto")

}