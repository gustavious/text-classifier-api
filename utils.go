package main

import (
	"math"
	"regexp"
	"strings"
)

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