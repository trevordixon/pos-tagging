package main

import (
	"./lib"
)

func GetTrainingData() (transition lib.ProbMatrix, emission lib.ProbMatrix, states *lib.ProbMap) {
	transition = make(lib.ProbMatrix)
	emission = make(lib.ProbMatrix)
	states = lib.NewProbMap()

	stream := lib.WordStream("data/allTraining.txt")
	prevWord := (<-stream).Value
	for word := range stream {
		transition.Observe(prevWord, word.Part)
		emission.Observe(word.Part, word.Value)
		states.Observe(word.Part)

		prevWord = word.Part
	}

	return
}

func main() {
	transition, emission, states := GetTrainingData()
	lib.Viterbi(states, nil, transition, emission)
}
