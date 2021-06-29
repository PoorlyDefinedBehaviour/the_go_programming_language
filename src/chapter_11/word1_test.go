package word1

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "detartrated",
			expected: true,
		},
		{
			input:    "palindrome",
			expected: false,
		},
		{
			input:    "été",
			expected: true,
		},
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
	}

	for _, tt := range tests {
		actual := IsPalindrome(tt.input)

		assert.Equal(t, tt.expected, actual)
	}
}
