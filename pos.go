package main

import (
	"./wordstream"
	"fmt"
	"strings"
)

type Suffixes map[string]WordCounts

func (s Suffixes) Observe(context string, word string) {
	s[context].Observe(word)
}

type WordCounts struct {
	Counts map[string]int
	Total  int
}

func (wc WordCounts) Observe(word string) {
	if wc.Counts == nil {
		wc.Counts = make(map[string]int)
	}

	wc.Counts[word]++
	wc.Total++
}

func (wc WordCounts) Prob(word string) float64 {
	return float64(wc.Counts[word]) / float64(wc.Total)
}

func NewWordcounts() (wc WordCounts) {
	wc = WordCounts{}
	wc.Counts = make(map[string]int)
	return
}

func GetPartsAndWords() (parts Suffixes, words Suffixes) {
	parts = make(Suffixes)
	words = make(Suffixes)

	context := []string{"", ""}
	for word := range wordstream.WordStream("data/allTraining.txt") {
		c := strings.Join(context, " ")

		parts.Observe(c, word.Part)
		words.Observe(word.Part, word.Value)

		context[0] = context[1]
		context[1] = word.Part
	}

	return
}

func main() {
	parts, words := GetPartsAndWords()
	fmt.Println(parts, words)

	// context := []string{"", ""}
	// for i := 0; i < 50; i++ {
	// 	c := strings.Join(context, " ")

	// 	fmt.Print(nextWord + " ")

	// 	context[0] = context[1]
	// 	context[1] = nextWord
	// }
}
