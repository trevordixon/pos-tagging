package lib

import (
	"bufio"
	"os"
	"strings"
)

type Word struct {
	Value string
	Part  string
}

func streamWords(path string, max int, stream chan Word) {
	file, err := os.Open(path)

	if err != nil {
		panic("Couldn't open " + path)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for i := 0; scanner.Scan(); i++ {
		if max > 0 && i >= max {
			break
		}

		d := strings.Split(scanner.Text(), "_")
		stream <- Word{d[0], d[1]}
	}

	close(stream)
}

func WordStream(file string, max int) (stream chan Word) {
	stream = make(chan Word, 100)
	go streamWords(file, max, stream)
	return
}
