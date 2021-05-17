package main

import (
	"image"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

/*
Functions may return multiple values
*/

// functions that return multiple values, return a
// tuple in fact
func f() (int, int) {
	return 0, 1
}

// the return values can be named
// for documentation purposes
func Size(rect image.Rectangle) (width, height int) {
	return 0, 0
}

func Split(path string) (directory, file string) {
	return "", ""
}

func HourMinSec(t time.Time) (hour, minute, second int) {
	return 0, 0, 0
}

// In a function with named results, the operands of a return
// statement may be omitted. This is called a bare return.
func CountWordsAndImages(url string) (words, images int, err error) {
	response, err := http.Get(url)
	if err != nil {
		return // equivalent to return words, images, err
	}

	defer response.Body.Close()

	document, err := html.Parse(response.Body)
	if err != nil {
		return // equivalent to return words, images, err
	}

	words, images = countWordsAndImages(document)
	return // equivalent to return words, images, err
}

func countWordsAndImages(node *html.Node) (words, images int) {
	return
}

func main() {
	False, True := f()
}
