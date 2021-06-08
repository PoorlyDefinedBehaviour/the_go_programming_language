package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

func walkDirectory(directory string, fileSizes chan<- int64) {
	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirectory := filepath.Join(directory, entry.Name())
			walkDirectory(subdirectory, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// This code is cursed.
func main() {
	verbose := flag.Bool("v", false, "show progress")

	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	// Why use a goroutine here?
	// Probably just a bad channels example.
	go func() {
		for _, root := range roots {
			root := root
			walkDirectory(root, fileSizes)
		}

		close(fileSizes)
	}()

	var tick <-chan time.Time = nil

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles int64 = 0
	var nbytes int64 = 0

	// why is this even in the language?
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes += size
		case <-tick:
		}
	}

	fmt.Printf("%d files %1.f GB\n", nfiles, float64(nbytes)/1e9)
}
