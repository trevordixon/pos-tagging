package wordstream

type ProbMatrix map[string]ProbMap
type ProbMap map[string]float64

func Viterbi(states ProbMap, observations []string, startProbability ProbMap,
	transitionProbability ProbMatrix, emissionProbability ProbMatrix) {

	for _ = range states {

	}
}
