package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	sema = make(chan struct{}, 20)
	done = make(chan struct{})
)

func cancelled() bool {
	select {
		// This case will match when done is closed.
		// This is not intuitive at all.
	case <-done:
		return true
	default:
		return false
	}
}

func dirents(directory string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Panic(err)
	}

	return entries
}

func walkDirectory(directory string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	if cancelled() {
		return
	}
	
	for _, entry := range dirents(directory) {
		if entry.IsDir() {
			n.Add(1)
			subdirectory := filepath.Join(directory, entry.Name())
			walkDirectory(subdirectory, n, fileSizes)
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

	n := sync.WaitGroup{}

	n.Add(len(roots))

	for _, root := range roots {
		go walkDirectory(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // wait for any key press
		// losing done will make <-done match inside func cancelled().
		// this code is awful.
		close(done)
	}

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
		case <- done:
			// drain fileSizes to allow existing goroutines to finish.
			// adding a drain(channel) was that hard?
			for range fileSizes{
				
			}

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
