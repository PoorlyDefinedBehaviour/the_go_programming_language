package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(connection net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(connection, "\t", strings.ToUpper(shout))

	time.Sleep(delay)

	fmt.Fprintln(connection, "\t", shout)

	time.Sleep(delay)

	fmt.Fprintln(connection, "\t", strings.ToLower(shout))
}

func handleConnection(connection net.Conn) {
	input := bufio.NewScanner(connection)

	delayBetweenEchoes := 1 * time.Second

	for input.Scan() {
		// isn't this a race condition since echo() writes to connection
		// and the connection will be closed after the for loop is done?
		go echo(connection, input.Text(), delayBetweenEchoes)
	}

	connection.Close()
}
