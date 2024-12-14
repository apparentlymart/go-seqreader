package seqreader_test

import (
	"fmt"
	"os"
	"slices"

	"github.com/apparentlymart/go-seqreader/seqreader"
)

func ExampleItemsFromSliceReader() {
	seq := slices.Values([][]byte{[]byte("hello "), []byte("world")})
	sliceR := seqreader.FromSeq(seq)
	r := seqreader.ItemsFromSliceReader(sliceR)

	for {
		b, err := r.ReadSeq()
		if err == seqreader.ErrEndOfSeq {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read: %s\n", err)
			break
		}
		fmt.Printf("byte: 0x%02x (%c)\n", b, b)
	}

	// Output:
	// byte: 0x68 (h)
	// byte: 0x65 (e)
	// byte: 0x6c (l)
	// byte: 0x6c (l)
	// byte: 0x6f (o)
	// byte: 0x20 ( )
	// byte: 0x77 (w)
	// byte: 0x6f (o)
	// byte: 0x72 (r)
	// byte: 0x6c (l)
	// byte: 0x64 (d)
}
