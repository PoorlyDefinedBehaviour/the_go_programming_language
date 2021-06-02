package main

import (
	"io"
	"log"
	"net"
	"os"
)

func must(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Panic(err)
	}

	defer connection.Close()

	_, err = io.Copy(os.Stdout, connection)
	must(err)

	_, err = io.Copy(connection, os.Stdin)
	must(err)
}
