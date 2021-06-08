package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
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

func main() {
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

	var nfiles int64 = 0
	var nbytes int64 = 0

	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	fmt.Printf("%d files %1.f GB\n", nfiles, float64(nbytes)/1e9)
}
