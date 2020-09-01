package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"message"`
}

// ProcessMessage classifies a given message
func getAnswerFromCategory(category string) string {
	switch category {
	case greeting:
		return "¡Bienvenido! Nos alegra verte de por aquí ¿En que te podemos ayudar?"
	case liked:
		return "¡Nos alegra mucho que te haya gustado! Esperamos que vuelvas pronto"
	case order:
		return "¿Qué deseas ordenar?"
	case food:
		return "¿Qué comida deseas?"
	case pizza:
		return "Tu pizza estará lista en cuestion de minutos"
	case hamburger:
		return "Marchando una hamburguesa"
	case salad:
		return "Perfecto, pronto tu ensalada estará lista"
	case soda:
		return "Una soda, entendido"
	}
	return "No he entendido eso ultimo"
}

// OptionsPreflight Handle cors for preflight OPTIONS
func OptionsPreflight(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
}


// ProcessMessage classifies a given message
func ProcessMessage(w http.ResponseWriter, r *http.Request) {
	// Handle cors for preflight OPTIONS
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)

	results := model.classify(msg.Text)
	category := ""

	maxWeight := 0.0
	for name, weight := range results {
		if  weight > maxWeight {
			maxWeight = weight
			category = name
		}
	}

	fmt.Println("Key word weights: ", results)
	respondwithJSON(w, http.StatusCreated, map[string]string{
		"awswer": getAnswerFromCategory(category),
		"category": category,
	})
}
