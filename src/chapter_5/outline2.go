package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Using a higher order function to reuse behaviour
func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, post)
	}

	if post != nil {
		post(node)
	}
}

var depth = 0

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		bytes, err := httpGet(url)
		if err != nil {
			panic(err)
		}

		doc, err := html.Parse(strings.NewReader(string(bytes)))
		if err != nil {
			panic(err)
		}

		forEachNode(doc, startElement, endElement)
	}
}

func httpGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
