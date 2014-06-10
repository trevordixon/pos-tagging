package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var counts map[string]string

func count() {
	file, err := os.Open("data/allTraining.txt")

	if err != nil {
		fmt.Errorf("Couldn't open allTraining.txt")
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		d := strings.Split(scanner.Text(), "_")
		// word is d[0], part of speech is d[1]
		fmt.Println(d)
	}
}

func main() {
	count()
}
