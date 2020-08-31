package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ProcessMessage classifies a given message
func ProcessMessage(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Println(post)

	results := model.classify(post.Message)
	class := ""

	maxWeight := 0.0
	for name, weight := range results {
		if  weight > maxWeight {
			maxWeight = weight
			class = name
		}
	}

	fmt.Println("Key word weights: ", results)
	fmt.Println("Classification: ", class)

	respondwithJSON(w, http.StatusCreated, map[string]string{
		"message": "successfully created",
		"Classification": class,
	})
}