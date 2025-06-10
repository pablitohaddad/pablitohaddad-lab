package auxiliar

import (
	"fmt"
	"strings"
)

func ExtractJSON(rawResp string) (string, error) {
	// Procura o início de um array JSON ou de um objeto JSON
	idxStartArray := strings.Index(rawResp, "[")
	idxStartObj := strings.Index(rawResp, "{")

	var start int

	// Descobre qual vem primeiro, ou se não existem
	if idxStartArray != -1 && (idxStartArray < idxStartObj || idxStartObj == -1) {
		start = idxStartArray
	} else if idxStartObj != -1 {
		start = idxStartObj
	} else {
		return "", fmt.Errorf("nenhum início de JSON ('[' ou '{') encontrado na resposta")
	}

	// Procura o fim correspondente
	var end int
	if start == idxStartArray {
		end = strings.LastIndex(rawResp, "]")
	} else {
		end = strings.LastIndex(rawResp, "}")
	}
	
	if end == -1 || end < start {
		return "", fmt.Errorf("nenhum fim de JSON correspondente (']' ou '}') encontrado na resposta")
	}

	// Retorna a fatia da string que contém apenas o JSON
	return rawResp[start : end+1], nil
}