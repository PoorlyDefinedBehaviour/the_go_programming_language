package word2

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true}, {"palindrome", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	// Use b.N when system under test input is static
	// and fixed iterations when it is not.
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// The third kind of function treated specially by go test if an example function,
// one whose name starts with Example. It has neither parameters not results.
// Here's an example function for IsPalindrome:
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

// A good example can be a more succinct or intuitive way to convey the behavior of
// a library function than its prose description.
// Unlike examples within comments, example functions are real Go code,
// subject to compile-time checking, so they don't become stale as
// the code evolves.
//
// Based on the suffix of the Example function, the web-based
// documentation server godoc associates example functions with
// the function or package they exemplify, so ExampleIsPalindrome
// would be shown with the documentation for the IsPalindrome function,
// and an example function called just Example would be associated
// with the word package as a whole.
//
// Example functions are executable tests run by go test. If the example
// function contains a final // Output: comment like the one in ExampleIsPalindrome,
// the test driver will execute the functionm and check that what it printed to its
// standard output matches the text within the comment(kinda like elixir doc tests https://elixir-lang.org/getting-started/mix-otp/docs-tests-and-with.html#doctests).
