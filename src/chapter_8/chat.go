package main

import (
	"bufio"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case client := <-entering:
			clients[client] = true
		case client := <-leaving:
			delete(clients, client)
			close(client)
		}
	}
}

func clientWritter(connection net.Conn, channel <-chan string) {
	for messages := range channel {
		fmt.Fprintln(connection, message)
	}
}

func handleConnection(connection net.Conn) {
	channel := make(chan string)

	go clientWriter(connection, channel)

	who := connection.RemoteAddr().String()

	channel <- "You are " + who

	messages <- who + " has arrived"

	entering <- channel

	input := bufio.NewScanner(connection)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- channel

	messages <- who + " has left"

	connection.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConnection(connection)
	}
}
