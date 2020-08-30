package main

import "fmt"

// The string values of the different classes
const (
	liked    = "liked"
	disliked = "disliked"
	greeting = "greeting"
	food = "food"
	order = "order"
	hamburger = "hamburger"
	soda = "soda"
	salad = "salad"
)

/*
 * Classifier
 */

// wordFrequency stores frequency of words.
type wordFrequency struct {
	word    string
	counter map[string]int
}

// classifier can be trained and used to categorize objects
// Attributes:
//	dataset: map each class with a list of  sentences from training data
//		map[string][]string{
//			"liked": []string{
//				"The restaurant is excellent",
//				"I really love this restaurant",
//			},
//			"disliked": []string{
//				"Their food is awful",
//			}
//
//		}
//	words: map each word with their frequency
//		map[string]wordFrequency{
//			"restaurant": wordFrequency{
//				word: "restaurant"
//				counter: map[string]int{
//					"liked": 2
//					"disliked": 0
//				}
// 			}
//		}
type classifier struct {
	dataset map[string][]string
	words   map[string]wordFrequency
}

// newClassifier returns a new classifier with empty dataset and words
func newClassifier() *classifier {
	c := new(classifier)
	c.dataset = map[string][]string{
		liked:    []string{},
		disliked: []string{},
	}
	c.words = map[string]wordFrequency{}

	return c
}

// train populates a classifier's dataset and words with input dataset map
// Sample dataset: map[string]string{
//	"The restaurant is excellent": "Positive",
//	"I really love this restaurant": "Positive",
//	"Their food is awful": "Negative",
//}
func (c *classifier) train(dataset map[string]string) {
	// fmt.Printf("> Dataset ", dataset)

	for sentence, class := range dataset {
		c.addSentence(sentence, class)
		words := tokenize(sentence)
		for _, w := range words {
			c.addWord(w, class)
		}
	}
	// TODO: Output to a file
	fmt.Println("Classifier Dataset", c.dataset);
}

// classify return the probabilities of a sentence being each class
func (c classifier) classify(sentence string) map[string]float64 {
	words := tokenize(sentence)
	likedProb := c.probability(words, liked)
	dislikedProb := c.probability(words, disliked)
	greetingProb := c.probability(words, greeting)
	foodProb := c.probability(words, food)
	orderProb := c.probability(words, order)
	hamburgerProb := c.probability(words, hamburger)
	saladProb := c.probability(words, salad)
	sodaProb := c.probability(words, soda)

	return map[string]float64{
		liked: likedProb,
		disliked: dislikedProb,
		greeting: greetingProb,
		food: foodProb,
		order: orderProb,
		hamburger: hamburgerProb,
		salad: saladProb,
		soda: sodaProb,
	}
}

// addSentence adds a sentence and its class to a classifier's dataset map
func (c *classifier) addSentence(sentence, class string) {
	c.dataset[class] = append(c.dataset[class], sentence)
}

// addSentence adds a word to a classifier's words map and update its frequency
func (c *classifier) addWord(word, class string) {
	wf, ok := c.words[word]
	if !ok {
		wf = wordFrequency{word: word, counter: map[string]int{
			liked:    0,
			disliked: 0,
			greeting: 0,
			food: 0,
			order: 0,
			hamburger: 0,
			soda: 0,
			salad: 0,
		}}
	}
	wf.counter[class]++
	c.words[word] = wf
	// fmt.Println("wf y ok ", wf, ok)
}


// priorProb returns the prior probability of each class of the classifier
// This probability is determined purely by the training dataset
func (c classifier) priorProb(class string) float64 {
	return float64(len(c.dataset[class])) / float64(
		len(c.dataset[liked])+len(c.dataset[disliked])+len(c.dataset[disliked])+len(c.dataset[greeting])+len(c.dataset[food])+
			len(c.dataset[order])+len(c.dataset[hamburger])+len(c.dataset[soda])+len(c.dataset[salad]),
	)
}

// totalWordCount returns the word count of a class (duplicated also count)
// If class provided is not liked or disliked, it returns
// the total word count in dataset.
func (c classifier) totalWordCount(class string) int {
	likeCount := 0
	dislikeCount := 0
	greetingCount := 0
	foodCount := 0
	orderCount := 0
	hamburgerCount := 0
	saladCount := 0
	sodaCount := 0
	for _, wf := range c.words {
		likeCount += wf.counter[liked]
		dislikeCount += wf.counter[disliked]
		greetingCount += wf.counter[greeting]
		foodCount += wf.counter[food]
		orderCount += wf.counter[order]
		hamburgerCount += wf.counter[hamburger]
		saladCount += wf.counter[salad]
		sodaCount += wf.counter[soda]
	}
	if class == liked {
		return likeCount
	} else if class == disliked {
		return dislikeCount
	} else if class == greeting {
		return greetingCount
	} else if class == food {
		return foodCount
	} else if class == order {
		return orderCount
	} else if class == hamburger {
		return hamburgerCount
	} else if class == salad {
		return saladCount
	} else if class == soda {
		return sodaCount
	} else {
		return likeCount + dislikeCount + greetingCount + foodCount + orderCount + hamburgerCount +
			saladCount + sodaCount
	}
}

// totalDistinctWordCount returns the number of distinct words in dataset
func (c classifier) totalDistinctWordCount() int {
	likeCount := 0
	dislikeCount := 0
	greetingCount := 0
	foodCount := 0
	orderCount := 0
	hamburgerCount := 0
	saladCount := 0
	sodaCount := 0

	for _, wf := range c.words {
		likeCount += zeroOneTransform(wf.counter[liked])
		dislikeCount += zeroOneTransform(wf.counter[disliked])
		greetingCount += zeroOneTransform(wf.counter[greeting])
		foodCount += zeroOneTransform(wf.counter[food])
		orderCount += zeroOneTransform(wf.counter[order])
		hamburgerCount += zeroOneTransform(wf.counter[hamburger])
		saladCount += zeroOneTransform(wf.counter[salad])
		sodaCount += zeroOneTransform(wf.counter[soda])
	}
	return likeCount + dislikeCount + greetingCount + foodCount + orderCount + hamburgerCount + saladCount + sodaCount
}

// probability retuns the probability of a list of words being in a class
func (c classifier) probability(words []string, class string) float64 {
	prob := c.priorProb(class)
	for _, w := range words {
		count := 0
		if wf, ok := c.words[w]; ok {
			count = wf.counter[class]
		}
		prob *= float64(count + 1) / float64(c.totalWordCount(class) + c.totalDistinctWordCount())
	}
	for _, w := range words {
		count := 0
		if wf, ok := c.words[w]; ok {
			count += wf.counter[liked] + wf.counter[disliked] + wf.counter[greeting] + wf.counter[food] +
							wf.counter[order] + wf.counter[hamburger] + wf.counter[salad] + wf.counter[soda]
		}
		prob /= float64(count + 1) / float64(c.totalWordCount("") + c.totalDistinctWordCount())
	}
	return prob
}
