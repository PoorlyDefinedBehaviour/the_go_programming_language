package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		fileContents, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "dup_3 %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(fileContents), "\n") {
			if _, ok := counts[line]; ok {
				counts[line] += 1
			} else {
				counts[line] = 1
			}
		}

		for line, count := range counts {
			if count > 1 {
				fmt.Printf("%d\t%s\n", count, line)
			}
		}
	}
}
