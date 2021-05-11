package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
		return
	}

	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup_2: %v\n", err)
			continue
		}

		countLines(file, counts)
		file.Close()
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		line := input.Text()
		if _, ok := counts[line]; ok {
			counts[line] += 1
		} else {
			counts[line] = 1
		}
	}
}
