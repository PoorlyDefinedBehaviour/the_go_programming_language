package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*
golang.org/x/net/html is a non-standard package, the golang.org/x/... repositories
hold packages design and maintained by the Go team for applications such as
networking, internationalized text processing, mobile platforms, image manipulation,
cryptography, and developers tool. The packages are not in the standard library
because they're still under development or because they're rarely needed by the majority
of Go programmers.
*/

func main() {
	for _, url := range os.Args[1:] {
		htmlString, err := httpGetString(url)
		if err != nil {
			panic(err)
		}

		doc, err := html.Parse(strings.NewReader(htmlString))
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}

		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	}
}

func httpGetString(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("getting %s: %s", url, response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func isAnchorTag(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "a"
}

func visit(links []string, node *html.Node) []string {
	if isAnchorTag(node) {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				links = append(links, attribute.Val)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = visit(links, child)
	}

	return links
}
