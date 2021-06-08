package main

import (
	"fmt"
	"log"
	"os"
)

func extract(url string) ([]string, error) {
	return make([]string, 0), nil
}

func crawl(tokens chan struct{}, url string) []string {
	fmt.Println(url)

	list, err := extract(url)
	if err != nil {
		log.Panic(err)
	}

	return list
}

func main() {
	// problem:
	// this is too parallel, could use all available file descriptors
	worklist := make(chan []string)

	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if seen[link] {
				continue
			}

			seen[link] = true

			link := link
			go func() {
				worklist <- crawl(link)
			}()
		}
	}
}
