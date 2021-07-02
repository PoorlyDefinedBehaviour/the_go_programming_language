package echo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_echo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		newline   bool
		separator string
		args      []string
		expected  string
	}{
		{
			newline:   true,
			separator: "",
			args:      []string{},
			expected:  "",
		},
	}
	for _, tt := range tests {
		// Mock dependency:
		// echo() writes the result to out.
		// This works because the test is in the same package
		// as the system under test and
		// out is a global variable.
		out = new(bytes.Buffer)

		err := echo(tt.newline, tt.separator, tt.args)

		assert.Nil(t, err)

		actual := out.(*bytes.Buffer).String()

		assert.Equal(t, tt.expected, actual)
	}
}
