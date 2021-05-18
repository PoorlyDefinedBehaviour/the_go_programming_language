package main

// nil can be used as a method receiver

// IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
// Sum(empty list) is 0
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

type Values map[string][]string

func (values Values) f() {

}

func main() {
	Values(nil).f() // ok
	//nil.f() // compile error
}
