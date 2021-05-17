package main

/*
A function for which failure is an expected behaviour returns an additional result,
conventionally the last one. If the failure has only one possible cause,
the result is a boolean, usually called ok, as in this example of a cache
lookup that always succeeds unless there was no entry for that key:
*/
type CacheOps interface {
	get(key string) (interface{}, bool)
}

type Cache struct {
}

/*
??? Let's be honest, a Result/Either monad + pattern matching is a better and more flexible
solution.
*/
func (c Cache) get(key string) (interface{}, bool) {
	return 1, true // returns false if key is not in the cache
}

/*
The built-in error type is an interface type. It can be nil.
A nil error implies success and non-nill implies failure and
that a non-nil error has an error message string which can be obtained
by calling its Error method or print by calling fmt.Println(err)

Usually when a function returns a non-nill error, its other results
are undefined and should be ignored. However, a few functions may return
partial results in error cases. For example, if an error occurs while
reading from a file, a call to Read returns the number of bytes it was
able to read and an error value describing the problem.

??? The book says that this problem is solved through documentation
which is funny because it could be solved by the compiler. There's nothing
stopping the caller of a function from using the undefined value.

??? Is go error handling actually any good? Since the error can be appear in any order in the result
list, it would probably be awkward to implement an error propagation opperator such as rust's ?.

Meanwhile, if go just had a Result monad, calling flatMap would be enough.

Ex:

value := functionThatMayReturnError()?

is equivalent to

value, err := functionThatMayReturnError()
if err != nil {
	return nil, err
}
*/

func main() {

}
