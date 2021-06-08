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

	tokens <- struct{}{} // acquire token

	list, err := extract(url)

	<-tokens // release token

	if err != nil {
		log.Panic(err)
	}

	return list
}

func main() {
	// tokens is a counting semaphore used to
	// enforce a limit of 20 concurrent requests.
	tokens := make(chan struct{}, 20)

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
				worklist <- crawl(tokens, link)
			}()
		}
	}
}
