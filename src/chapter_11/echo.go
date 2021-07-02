package echo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// This global variable exists so we can
// swap the output implementation during tests.
// DI is another(probably better) option.
var out io.Writer = os.Stdout

func echo(newline bool, separator string, args []string) error {
	_, err := fmt.Fprint(out, strings.Join(args, separator))
	if err != nil {
		return errors.WithStack(err)
	}

	if newline {
		_, err := fmt.Fprintln(out)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
