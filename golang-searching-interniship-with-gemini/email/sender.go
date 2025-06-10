package email

import (
	"encoding/json"
	"log"
	"os"
	bodyhtml "using-gemini-api/email/body_html"
	"using-gemini-api/models"

	"gopkg.in/gomail.v2"
)


func SendEmail(respostaGemini string){

	// Passo a resposta do gemini para a struct vaga
	var vagas []models.Vaga
	err := json.Unmarshal([]byte(respostaGemini), &vagas)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do JSON: %v", err)
	}

	// Construo o html resposta
	htmlBody := bodyhtml.BuildHTMLEmail(vagas)

	// Envio o email
	d := gomail.NewDialer("smtp.gmail.com", 587, "pablohaddad73@gmail.com", os.Getenv("GOOGLE_API_KEY"))
	mailForMe := gomail.NewMessage()
	mailForMe.SetHeader("From", "pablohaddad73@gmail.com")
	mailForMe.SetHeader("To", "pablitohaddad@gmail.com")
	mailForMe.SetHeader("Subject", "Vagas de estágio de hoje! - Pablo Haddad")
	mailForMe.SetBody("text/html", htmlBody)

	if err := d.DialAndSend(mailForMe); err != nil {
		panic(err)
	}
}

	


