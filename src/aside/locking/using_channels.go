package main

type Connection struct {
}

type ConnectionPool struct {
	Connections chan Connection
}

func (connectionPool *ConnectionPool) Acquire() Connection {
	connection := <-connectionPool.Connections

	return connection
}

func (connectionPool *ConnectionPool) Release(connection Connection) {
	connectionPool.Connections <- connection
}

func main() {

}
