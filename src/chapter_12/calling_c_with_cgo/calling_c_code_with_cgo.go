package main

// cgo is a tool that create Go bindings for C functions.
// Such tools are called foreign-function interfaces(FFI),
// and cgo is not the only one for Go programs.
// SWIG(swig.org) is another.
//
// During preprocessing, cgo generates a temporary package
// that contains Go declarations corresponding to all the C
// functions and types used by the file.
// The cgo tool discovers these types by invoking the C
// compiler in a special way on the contents of the comment
// that precedes the import declaration.
// The comment may also contain #cgo directives that specify extra
// options to the C toolchain.
// The CFLAGS and LDFLAGS values contribute extra arguments
// to the compiler and linker commands so that hey can locate the bzlib.h
// header file and the libbz2.a archive library.

// This import delcaration is special. There is no package C,
// but this import causes go build to preprocess the file
// using the cgo tool before the Go compiler sees it.
/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
int bz2compress(bz_stream *s, int action, char* in, unsigned* inlen, char* out, unsigned* outlen);
*/
import "C"
import (
	"io"
	"unsafe"

	"github.com/pkg/errors"
)

type writer struct {
	w      io.Writer
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)

	w := &writer{w: out, stream: new(C.bz_stream)}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	if w.stream == nil {
		panic("closed")
	}

	total := 0

	for len(data) > 0 {
		inlen := C.uint(len(data))
		outlen := C.uint(cap(w.outbuf))

		C.bz2compress(
			w.stream,
			C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])),
			(*C.char)(unsafe.Pointer(&w.outbuf)),
		)

		total := int(inlen)

		data = data[inlen:]

		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, errors.WithStack(err)
		}
	}

	return total, nil
}

func (w *writer) Close() error {
	if w.stream == nil {
		panic("closed")
	}

	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		w.stream = nil
	}()

	for {
		inlen := C.uint(0)
		outlen := C.uint(cap(w.outbuf))

		r := C.bz2compress(
			w.stream,
			C.BZ_FINISH,
			nil,
			&inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)),
		)

		if err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return errors.WithStack(err)
		}

		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}

func main() {

}
