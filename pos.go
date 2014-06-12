package main

import (
	"./lib"
	"./pb"
	"log"
	"time"
)

func GetTrainingData() (transition lib.ProbMatrix, emission lib.ProbMatrix, states *lib.ProbMap) {
	transition = make(lib.ProbMatrix)
	emission = make(lib.ProbMatrix)
	states = lib.NewProbMap()

	stream := lib.WordStream("data/allTraining.txt", 0)
	prevWord := (<-stream).Value
	for word := range stream {
		transition.Observe(prevWord, word.Part)
		emission.Observe(word.Part, word.Value)
		states.Observe(word.Part)

		prevWord = word.Part
	}

	return
}

func reportProgress(total int, progress chan bool) {
	bar := pb.New(total)
	bar.SetRefreshRate(time.Millisecond * 60)
	bar.Start()

	for _ = range progress {
		bar.Increment()
	}
}

func main() {
	log.Println("Loading training data")

	transition, emission, states := GetTrainingData()

	words := []string{}
	parts := []string{}
	for word := range lib.WordStream("data/devtest.txt", 1000) {
		words = append(words, word.Value)
		parts = append(parts, word.Part)
	}

	log.Println("Tagging ", len(words), " words")

	progress := make(chan bool)
	go reportProgress(len(words), progress)

	class := lib.Viterbi(states, words, transition, emission, progress)

	correct := 0
	for i, part := range class {
		if part == parts[i] {
			correct++
		}
	}

	pctCorrect := (float64(correct) / float64(len(words))) * 100
	log.Println("Percent correct: ", pctCorrect)
	// log.Println("Tags: ", class)
}
