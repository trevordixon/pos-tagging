package main

import (
	"./lib"
)

func GetPartsAndWords() (parts lib.ProbMatrix, words lib.ProbMatrix, states *lib.ProbMap) {
	parts = make(lib.ProbMatrix)
	words = make(lib.ProbMatrix)
	states = lib.NewProbMap()

	stream := lib.WordStream("data/allTraining.txt")
	prevWord := (<-stream).Value
	for word := range stream {
		parts.Observe(prevWord, word.Part)
		words.Observe(word.Part, word.Value)
		states.Observe(word.Part)

		prevWord = word.Part
	}

	return
}

func main() {
	parts, words, states := GetPartsAndWords()

	lib.Viterbi(states, nil, parts, words)
}
