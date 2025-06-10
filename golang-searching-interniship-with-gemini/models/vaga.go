package models

type Vaga struct {
	Titulo      string   `json:"titulo_vaga"`
	Empresa     string   `json:"empresa"`
	Cargo       string   `json:"cargo"`
	Tecnologias []string `json:"tecnologias_principais"`
	Local       string   `json:"local"`
	Data        string   `json:"data_publicacao"`
	URL         string   `json:"url_vaga"`
	IsEstagio   bool     `json:"is_estagio"`
}