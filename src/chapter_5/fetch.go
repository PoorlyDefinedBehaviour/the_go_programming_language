package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

// Requests URL, saves http response to a local file.
// The file name is the last component of the URL path.
func fetch(url string) (filename string, length int64, err error) {
	response, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer response.Body.Close()

	filename = path.Base(response.Request.URL.Path)
	if filename == "/" {
		filename = "index.html"
	}

	file, err := os.Create(filename)
	if err != nil {
		return "", 0, err
	}

	length, err = io.Copy(file, response.Body)

	if err != nil {
		file.Close()
		return "", 0, err
	}

	err = file.Close()
	if err != nil {
		return "", 0, err
	}

	return filename, length, err
}

func main() {

}
