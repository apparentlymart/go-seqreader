package seqreader_test

import (
	"fmt"
	"os"
	"slices"

	"github.com/apparentlymart/go-seqreader/seqreader"
)

func ExampleFromSeq() {
	seq := slices.Values([]string{"hello", "world"})
	r := seqreader.FromSeq(seq)

	for {
		s, err := r.ReadSeq()
		if err == seqreader.ErrEndOfSeq {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read: %s\n", err)
			break
		}
		fmt.Println(s)
	}

	// Output:
	// hello
	// world
}
