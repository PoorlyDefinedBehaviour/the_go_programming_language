package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"

	"github.com/pkg/errors"
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Fprintln(os.Stderr, "Input format = ", kind)

	err = jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func main() {
	fmt.Printf("\n\naaaaaaa  %+v\n\n", errors.WithStack(errors.New("oops")))
	//fmt.Printf("\n\naaaaaaa   %+v\n\n", errors.Wrap(errors.New("oops"), "test"))
	// if err := toJPEG(os.Stdin, os.Stdout); err != nil {
	// 	log.Panic(err)
	// }
}
