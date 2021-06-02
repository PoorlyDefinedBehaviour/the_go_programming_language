package main

import "io"

func writeString(writer io.Writer, s string) (int, error) {
	type StringWriter interface {
		WriteString(string) (int, error)
	}

	if stringWriter, ok := writer.(StringWriter); ok {
		return stringWriter.WriteString(s) // avoids copy
	}

	return writer.Write([]byte(s)) // allocates temporary copy
}

func main() {

}
