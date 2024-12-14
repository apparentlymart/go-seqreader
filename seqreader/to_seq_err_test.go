package seqreader_test

import (
	"fmt"
	"slices"

	"github.com/apparentlymart/go-seqreader/seqreader"
)

func ExampleToSeq() {
	// This is a SeqReader[string] that fails on its third read.
	r := SeqReaderThatFails()
	messages := seqreader.ToSeq(r)

	// Use the result of the Items method with something that expects
	// an infallible, finite iter.Seq.
	slice := slices.Collect(messages.Items())
	// After that operation is complete, call Err to find out if
	// the traversal ended early due to an error.
	if err := messages.Err(); err != nil {
		fmt.Printf("err = %q\n", err)
	}
	fmt.Printf("slice = %#v\n", slice)

	// Output:
	// err = "failed to ding the wotsit"
	// slice = []string{"hello", "world"}
}
