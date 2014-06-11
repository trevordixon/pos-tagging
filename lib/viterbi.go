package lib

import (
	"fmt"
)

func Viterbi(states *ProbMap, observations []string,
	transitionProbability ProbMatrix, emissionProbability ProbMatrix) {

	fmt.Println(emissionProbability["NNS"].Prob("dogs"))

	// for state := range states.Counts {
	// 	fmt.Println(transitionProbability[state].Prob("AV"))
	// 	fmt.Println(emissionProbability[state].Prob("dog"))
	// }
}
