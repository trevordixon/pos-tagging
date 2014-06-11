package main

import (
	"./lib"
	"fmt"
	"math/rand"
	"strings"
)

type Suffixes map[string]WordCounts
type WordCounts map[string]int

func (wc WordCounts) GetRandom() (word string) {
	sum := 0
	for _, count := range wc {
		sum += count
	}

	r := rand.Intn(sum)

	sum = 0
	for word, count := range wc {
		sum += count
		if r < sum {
			return word
		}
	}

	panic("Didn't find a random word.")
}

func GetSuffixes() (suffixes Suffixes) {
	suffixes = make(Suffixes)
	context := []string{"", ""}

	for word := range lib.WordStream("data/allTraining.txt") {
		c := strings.Join(context, " ")

		if suffixes[c] == nil {
			suffixes[c] = make(WordCounts)
		}
		suffixes[c][word.Value]++

		context[0] = context[1]
		context[1] = word.Value
	}

	return
}

func main() {
	suffixes := GetSuffixes()
	context := []string{"", ""}
	for i := 0; i < 100; i++ {
		c := strings.Join(context, " ")
		nextWord := suffixes[c].GetRandom()

		fmt.Printf("%s ", nextWord)

		context[0] = context[1]
		context[1] = nextWord
	}
}
