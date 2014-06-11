package lib

import (
	"fmt"
)

func Viterbi(states *ProbMap, observations []string, tr ProbMatrix, em ProbMatrix) {
	fmt.Println(em["NNS"].Prob("dogs"))

	// for state := range states.Counts {
	// 	fmt.Println(tr[state].Prob("AV"))
	// 	fmt.Println(em[state].Prob("dog"))
	// }
}
