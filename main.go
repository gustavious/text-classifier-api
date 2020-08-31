package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// returns a map of sentences with their initial class
func loadData(file string) map[string]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dataMap := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, "#")
		if len(data) != 2 {
			continue
		}
		sentence := data[0]
		dataMap[sentence] = data[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataMap
}

// Initialize a new classifier
var model = newClassifier()

func main() {
	// Initialize router
	routers()

	// train  model
	dataset := loadData("./chats")
	fmt.Println("Dataset: ", dataset)
	model.train(dataset)

	// Serve api
	fmt.Println("Server listening on port: ", 8005)
	http.ListenAndServe(":8005", Logger())
}
