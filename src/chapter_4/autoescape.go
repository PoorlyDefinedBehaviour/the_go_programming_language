package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t := template.Must(template.New("escape").Parse(`<p>A: {{.A}}</p><p>B: {{.B}}</p>`))

	type foo struct {
		A string        // untrusted plain text
		B template.HTML // trusted html
	}

	data := foo{
		A: "<b>Hello!</b>",
		B: "<b>Hello!</b>",
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
