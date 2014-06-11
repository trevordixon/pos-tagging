package wordstream

import (
	"bufio"
	"os"
	"strings"
)

type Word struct {
	Value string
	Part  string
}

func streamWords(path string, stream chan Word) {
	file, err := os.Open(path)

	if err != nil {
		panic("Couldn't open " + path)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		d := strings.Split(scanner.Text(), "_")
		stream <- Word{d[0], d[1]}
	}

	close(stream)
}

func WordStream(file string) (stream chan Word) {
	stream = make(chan Word, 100)
	go streamWords(file, stream)
	return
}
