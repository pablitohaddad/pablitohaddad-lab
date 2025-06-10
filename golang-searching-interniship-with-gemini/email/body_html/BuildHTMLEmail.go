package bodyhtml

import (
	"fmt"
	"strings"
	"using-gemini-api/models"
)

// BuildHTMLEmail constrói o corpo do e-mail em HTML a partir de uma lista de vagas.
func BuildHTMLEmail(vagas []models.Vaga) string {
	// Usamos um strings.Builder para construir a string de forma eficiente.
	var body strings.Builder

	// --- Início do Template HTML ---
	body.WriteString(`
		<!DOCTYPE html>
		<html lang="pt-br">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<style>
				body { font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f4f4f4; }
				.container { max-width: 600px; margin: 20px auto; background-color: #ffffff; padding: 20px; border-radius: 8px; }
				.header { text-align: center; padding-bottom: 20px; border-bottom: 1px solid #dddddd; }
				.job { border-bottom: 1px solid #eeeeee; padding: 20px 0; }
				.job-title { font-size: 20px; color: #333333; margin: 0; }
				.company { font-size: 16px; color: #555555; margin: 5px 0; }
				.details p { margin: 5px 0; color: #666666; }
				.tech { display: inline-block; background-color: #e0eafc; color: #3b5998; padding: 5px 10px; border-radius: 15px; font-size: 12px; margin-right: 5px; margin-top: 5px;}
				.button { display: block; width: 150px; margin: 20px auto; padding: 10px; background-color: #28a745; color: #ffffff; text-align: center; text-decoration: none; border-radius: 5px; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>🚀 Vagas de Estágio Encontradas Hoje!</h1>
					<p>Aqui está seu resumo diário de oportunidades.</p>
				</div>
	`)

	// Itera sobre cada vaga e cria uma seção para ela
	for _, vaga := range vagas {
		body.WriteString(fmt.Sprintf(`
			<div class="job">
				<h2 class="job-title">%s</h2>
				<h3 class="company">%s</h3>
				<div class="details">
					<p><strong>📍 Local:</strong> %s</p>
					<p><strong>🗓️ Publicado em:</strong> %s</p>
					<div>`,
			vaga.Titulo, vaga.Empresa, vaga.Local, vaga.Data))

		// Adiciona as "pílulas" de tecnologia
		for _, tech := range vaga.Tecnologias {
			body.WriteString(fmt.Sprintf(`<span class="tech">%s</span>`, tech))
		}
		
		// Fecha as divs e adiciona o botão de candidatura
		body.WriteString(fmt.Sprintf(`
				</div>
				<a href="%s" class="button" style="color: #ffffff;">Ver Vaga</a>
			</div>
		`, vaga.URL))
	}

	// --- Final do Template HTML ---
	body.WriteString(`
			</div>
		</body>
		</html>
	`)

	return body.String()
}