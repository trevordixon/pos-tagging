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

	for t, obs := range observations[1:] {
		V = append(V, FloatMap{})
		newPath := make(map[string][]string)

		for state := range states.Counts {
			maxVal := 0.0
			maxState := ""
			for state0 := range states.Counts {
				// Calculate the probablity
				calc := V[t][state0] * tr[state0].Prob(state) * em[state].Prob(obs)
				if calc > maxVal {
					maxVal = calc
					maxState = state0
				}
			}
			V[t+1][state] = maxVal
			newPath[state] = append(path[maxState], state)
		}
		path = newPath
	}

	maxVal := 0.0
	maxState := ""
	for state := range states.Counts {
		calc := V[len(observations)-1][state]
		if calc > maxVal {
			maxVal = calc
			maxState = state
		}
	}

	fmt.Println(maxVal)
	fmt.Println(path)

	for i, state := range path[maxState] {
		fmt.Println(i, state)
	}
}
