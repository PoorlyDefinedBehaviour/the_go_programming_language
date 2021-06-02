package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConnection(connection net.Conn) {
	// should't it be closed by who opened it since it is not async?
	defer connection.Close()

	for {
		_, err := io.WriteString(connection, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// blocks until a connection until there's an incoming connection
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// probably should be async
		handleConnection(connection)
	}
}
