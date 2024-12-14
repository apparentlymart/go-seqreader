package seqreader_test

import (
	"fmt"

	"github.com/apparentlymart/go-seqreader/seqreader"
)

func ExampleToSeq2() {
	// This is a SeqReader[string] that fails on reads three and onward.
	r := SeqReaderThatFails()
	for msg, err := range seqreader.ToSeq2(r) {
		if err != nil {
			fmt.Printf("error: %s\n", err)
			break // this sequence never stops returning errors, so must break
		}
		fmt.Printf("msg = %q\n", msg)
	}

	// Output:
	// msg = "hello"
	// msg = "world"
	// error: failed to ding the wotsit
}
