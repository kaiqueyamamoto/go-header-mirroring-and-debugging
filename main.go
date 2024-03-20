package main

import (
	"encoding/json"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	// Definindo o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Criando um mapa para armazenar os cabeçalhos
	headersMap := make(map[string]string)

	// Iterando sobre os cabeçalhos da requisição recebida e os adicionando ao mapa
	for name, values := range r.Header {
		// O cabeçalho pode ter múltiplos valores, então pegamos apenas o primeiro
		headersMap[name] = values[0]
	}

	// Convertendo o mapa de cabeçalhos para JSON
	jsonHeaders, err := json.Marshal(headersMap)
	if err != nil {
		http.Error(w, "Erro ao converter cabeçalhos para JSON", http.StatusInternalServerError)
		return
	}

	// Escrevendo o JSON de cabeçalhos na resposta
	w.Write(jsonHeaders)
}

func main() {
	// Definindo o manipulador da rota
	http.HandleFunc("/headers", headersHandler)

	// Iniciando o servidor na porta 8080
	println("Servidor iniciado na porta 2020")
	if err := http.ListenAndServe(":2020", nil); err != nil {
		panic(err)
	}
}
