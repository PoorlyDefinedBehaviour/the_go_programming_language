package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Panic(err)
	}

	defer connection.Close()

	mustCopy(os.Stdout, connection)
}

func mustCopy(destination io.Writer, source io.Reader) {
	if _, err := io.Copy(destination, source); err != nil {
		log.Panic(err)
	}
}
