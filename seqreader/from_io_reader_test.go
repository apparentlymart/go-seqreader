package seqreader_test

import (
	"bytes"
	"fmt"
	"os"

	"github.com/apparentlymart/go-seqreader/seqreader"
)

func ExampleFromIOReader() {
	buf := []byte(`hello, world of fallible sequences!`)
	ioR := bytes.NewReader(buf)
	r := seqreader.FromIOReader(ioR, func() []byte {
		fmt.Println("(allocating buffer)")
		return make([]byte, 4, 16)
	})

	for {
		chunk, err := r.ReadSeq()
		if err == seqreader.ErrEndOfSeq {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read: %s\n", err)
			break
		}
		fmt.Printf("chunk: %q\n", chunk)
	}

	// Output:
	// (allocating buffer)
	// chunk: "hell"
	// chunk: "o, w"
	// chunk: "orld"
	// chunk: " of "
	// (allocating buffer)
	// chunk: "fall"
	// chunk: "ible"
	// chunk: " seq"
	// chunk: "uenc"
	// (allocating buffer)
	// chunk: "es!"
}
