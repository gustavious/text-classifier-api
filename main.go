package main

import (
	"bufio"
	"log"
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

/*
func main() {
	// Initialize a new classifier
	model := newClassifier()
	dataset := loadData("./chats")
	model.train(dataset)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(" ")
		fmt.Print("En qué te podemos ayudar?: ")
		sentence, _ := reader.ReadString('\n')
		results := model.classify(sentence)
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
	}
}
 */