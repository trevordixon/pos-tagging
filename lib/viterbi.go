package lib

import (
	"fmt"
)

type FloatMap map[string]float64

func Viterbi(states *ProbMap, observations []string, tr ProbMatrix, em ProbMatrix) {
	V := []FloatMap{make(FloatMap)}
	path := make(map[string][]string)

	for state := range states.Counts {
		V[0][state] = states.Prob(state) * em[state].Prob(observations[0])
		path[state] = []string{state}
	}

	fmt.Println(V)
}
