package main

import (
	"log"
	"os"
	"sync"
)

func resizeImage(filename string) (string, error) {
	// pretend to resize
	return filename, nil
}

// could be concurrent
func makeThumbnails1(filenames []string) {
	for _, filename := range filenames {
		if _, err := resizeImage(filename); err != nil {
			log.Panic(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	doneChannel := make(chan struct{})

	for _, filename := range filenames {
		filename := filename

		go func() {
			_, err := resizeImage(filename)
			if err != nil {
				log.Panic(err)
			}

			doneChannel <- struct{}{}
		}()
	}

	for range filenames {
		<-doneChannel
	}
}

func makeThumbnails3(filenames []string) ([]string, error) {
	type Item struct {
		thumb string
		err   error
	}

	resizeResultChannel := make(chan Item, len(filenames))

	// Could have been:
	//
	// thumbs := Task.all(filenames.map(func(filename) Item {
	//	thumb, err := resizeImage(fileName)
	//	return Item{ thumb: tumb, err: err, }
	// }))
	//
	// You need to try really hard to write incorrect code that way.
	//
	for _, filename := range filenames {
		filename := filename

		go func() {
			thumb, err := resizeImage(filename)

			result := Item{
				thumb: thumb,
				err:   err,
			}

			resizeResultChannel <- result
		}()
	}

	thumbs := make([]string, 0, len(filenames))

	for range filenames {
		result := <-resizeResultChannel

		if result.err != nil {
			return nil, result.err
		}

		thumbs = append(thumbs, result.thumb)
	}

	return thumbs, nil
}

func makeThumbnails4(filenames <-chan string) int64 {
	sizes := make(chan int64)

	waitGroup := sync.WaitGroup{}

	for filename := range filenames {
		waitGroup.Add(1)

		filename := filename

		go func() {
			defer waitGroup.Done()

			thumb, err := resizeImage(filename)
			if err != nil {
				log.Println(err)
				return
			}

			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}()
	}

	go func() {
		waitGroup.Wait()
		close(sizes)
	}()

	// size.reduce(+)
	var total int64 = 0
	for size := range sizes {
		total += size
	}

	return total
}

func main() {

}
