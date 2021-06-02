package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	xmlDecoder := xml.NewDecoder(os.Stdin)

	stack := make([]string, 0)

	for {
		token, err := xmlDecoder.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panic(err)
		}

		switch token := token.(type) {
		case xml.StartElement:
			stack = append(stack, token.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), token)
			}
		}
	}
}

func containsAll(stack []string, values []string) bool {
	for len(values) <= len(stack) {
		if len(values) == 0 {
			return true
		}

		if stack[0] == values[0] {
			values = values[1:]
		}

		stack = stack[1:]
	}

	return false
}
