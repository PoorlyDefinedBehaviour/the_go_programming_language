package main

// error is just an interface type with a single method:
//
// type error interface {
// 	Error() string
// }

// The simplest way to create an error is by calling errors.New,
// which returns a new error for a given error message.
// The entire errors package is only a few lines long:
//
// package errors
//
// func New(text string) error {
//   // Returns an struct to make mutations to the text message impossible.
//   // Returns a pointer because errors with the same error message shouldn't be
//   // considered equal. Allocating them in the heap, makes they different when compared with ==.
//   // fmt.Println(errors.New("EOF") == errors.New("EOF")) -> false
//	 return  &errorString{text}
// }
//
// type errorString struct { text string }
//
// func(e *errorString) Error() string { return e.text }

func main() {

}
