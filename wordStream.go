package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Word struct {
	Value string
	Part  string
}

func streamWords(stream chan Word) {
	file, err := os.Open("data/allTraining.txt")

	if err != nil {
		fmt.Errorf("Couldn't open allTraining.txt")
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

func WordStream() (stream chan Word) {
	stream = make(chan Word)
	go streamWords(stream)
	return
}
