package lib

type ProbMatrix map[string]*ProbMap

func (s ProbMatrix) Observe(context string, word string) {
	if s[context] == nil {
		s[context] = NewProbMap()
	}
	s[context].Observe(word)
}

type ProbMap struct {
	Counts map[string]int
	Total  int
}

func (pm *ProbMap) Observe(word string) {
	if pm.Counts == nil {
		pm.Counts = make(map[string]int)
	}

	pm.Counts[word]++
	pm.Total++
}

func (pm *ProbMap) Prob(word string) float64 {
	return float64(pm.Counts[word]) / float64(pm.Total)
}

func NewProbMap() (pm *ProbMap) {
	pm = &ProbMap{Counts: make(map[string]int)}
	return
}
