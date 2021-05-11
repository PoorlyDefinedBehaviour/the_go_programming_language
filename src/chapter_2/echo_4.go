package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var omitTrailingNewline = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*omitTrailingNewline {
		fmt.Println()
	}
}
