package gemini

import (
	"context"
	"fmt"
	"log"

	// CORREÇÃO 1: Usando o nome do seu módulo para importar o pacote local.
	auxiliar "using-gemini-api/auxiliar"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AnalyzeText(apiKey, textToAnalyze string) (string, error) {
	// Conexão com o Gemini
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// Modelo e Instruções
	model := client.GenerativeModel("gemini-2.0-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("Você é um assistente de RH para tecnologia, focado em vagas de estágio no Brasil. Sua tarefa é analisar o texto de uma vaga e extrair as informações solicitadas no formato JSON. Responda apenas com o JSON."),
		},
	}
    model.GenerationConfig.ResponseMIMEType = "application/json"

	// Prompt para o programa
	prompt := fmt.Sprintf(`
        --- DESCRIÇÃO DA VAGA ---
        %s
        --- FIM DA DESCRIÇÃO ---

        Responda APENAS com um objeto JSON válido, com a seguinte estrutura:
        [
			{
          		"titulo_vaga": "O nome da vaga",
          		"empresa": "O nome da empresa",
          		"cargo": "A função principal (ex: Backend, Frontend, Full Stack)",
          		"tecnologias_principais": ["Uma", "lista", "de", "tecnologias"],
          		"local": "O local da vaga (ex: Remoto, São Paulo - SP)",
          		"data_publicacao": "A data em que a vaga foi postada",
          		"url_vaga": "O link para se inscrever na vaga",
          		"is_estagio": true
        	}
		]`, textToAnalyze)

	// Pegando a resposta do modelo
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("erro ao gerar conteúdo: %v", err)
	}
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("resposta inesperada ou vazia do Gemini")
	}


	// Extaindo o texto bruto da Gemini
	var rawTextFromGemini string
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			rawTextFromGemini += string(txt)
		}
	}
	if rawTextFromGemini == "" {
		return "", fmt.Errorf("a resposta do Gemini não continha texto")
	}

	// Limpando o json bruto
	jsonLimpo, err := auxiliar.ExtractJSON(rawTextFromGemini)
	if err != nil {
		return "", fmt.Errorf("erro ao extrair JSON da resposta: %w", err)
	}
	return jsonLimpo, nil
}