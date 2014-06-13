package lib

import (
	"math"
	"strings"
)

type FloatMap map[string]float64

func Viterbi(states *ProbMap, observations []string, tr ProbMatrix, em ProbMatrix, progress chan bool) []string {
	V := []FloatMap{make(FloatMap)}
	path := make(map[string][]string)

	for state := range states.Counts {
		emKey := strings.Split(state, " ")[1]

		V[0][state] = math.Log(states.Prob(state)) + math.Log(em[emKey].Prob(observations[0]))
		path[state] = []string{state}
	}

	for t, obs := range observations[1:] {
		V = append(V, FloatMap{})
		newPath := make(map[string][]string)

		for state := range states.Counts {
			emKey := strings.Split(state, " ")[1]
			maxVal := math.Inf(-1)
			maxState := ""
			for state0 := range states.Counts {
				// Calculate the probablity
				calc := V[t][state0] + math.Log(tr[state0].Prob(state)) + math.Log(em[emKey].Prob(obs))
				if calc > maxVal {
					maxVal = calc
					maxState = state0
				}
			}
			V[t+1][state] = maxVal
			newPath[state] = make([]string, len(path[maxState]))
			copy(newPath[state], path[maxState])
			newPath[state] = append(newPath[state], state)
		}
		path = newPath

		progress <- true
	}

	close(progress)

	maxVal := math.Inf(-1)
	maxState := ""
	for state := range states.Counts {
		calc := V[len(observations)-1][state]
		if calc > maxVal {
			maxVal = calc
			maxState = state
		}
	}

	return path[maxState]
}
