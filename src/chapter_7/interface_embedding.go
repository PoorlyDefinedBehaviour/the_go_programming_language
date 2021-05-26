package main

type Writer interface {
	Write(p []byte) (bytesWritten int, err error)
}

type Reader interface {
	Read(p []byte) (bytesRead int, err error)
}

type Closer interface {
	Close() error
}

// interface embedding
type ReadWriter interface {
	Reader
	Writer
}

// interface embedding
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

func main() {

}
