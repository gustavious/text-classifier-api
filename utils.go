package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}

// writeOutputFile writes to output file
func writeOutputFile(classifier *classifier) {
	// write the whole body at once
	out, _ := json.Marshal(classifier.dataset)

	err := ioutil.WriteFile("classifierRawOutputData.txt", out, 0644)
	if err != nil {
		panic(err)
	}
}

// stopwords are words which have very little meaning
var stopwords = map[string]struct{}{
	"la": struct{}{},
}

func isStopword(w string) bool {
	_, ok := stopwords[w]
	return ok
}

// tokenize create an array of words from a sentence
func tokenize(sentence string) []string {
	s := cleanup(sentence)
	words := strings.Fields(s)
	var tokens []string
	for _, w := range words {
		if !isStopword(w) {
			tokens = append(tokens, w)
		}
	}
	return tokens
}

// zeroOneTransform returns
//   0 if argument x = 0; 1 otherwise
func zeroOneTransform(x int) int {
	return int(math.Ceil(float64(x) / (float64(x) + 1.0)))
}

// cleanup remove non alphanumeric chars and lowercasize them
func cleanup(sentence string) string {
	re := regexp.MustCompile("[^a-zA-Z 0-9]+")
	return re.ReplaceAllString(strings.ToLower(sentence), "")
}