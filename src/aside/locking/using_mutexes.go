package main

import "sync"

type Connection struct {
}

type ConnectionPool struct {
	Lock        sync.Mutex
	Connections []Connection
}

func (connectionPool *ConnectionPool) Aquire() Connection {
	connectionPool.Lock.Lock()
	defer connectionPool.Lock.Unlock()

	// get a connectionm from the pool and return it

	return Connection{}
}

func (connectionPool *ConnectionPool) Release(connection Connection) {
	connectionPool.Lock.Lock()
	defer connectionPool.Lock.Unlock()

	// give connection back to the pool
}

func main() {

}
