package seqreader_test

import (
	"fmt"

	"github.com/apparentlymart/go-seqreader/seqreader"
)

func SeqReaderThatFails() seqreader.SeqReader[string] {
	ret := failSeqReader(0)
	return &ret
}

type failSeqReader int

func (r *failSeqReader) ReadSeq() (string, error) {
	state := *r
	*r++
	switch state {
	case 0:
		return "hello", nil
	case 1:
		return "world", nil
	default:
		return "", fmt.Errorf("failed to ding the wotsit")
	}
}

func SeqReaderThatSaysHello() seqreader.SeqReader[string] {
	ret := helloSeqReader(0)
	return &ret
}

type helloSeqReader int

func (r *helloSeqReader) ReadSeq() (string, error) {
	state := *r
	*r++
	switch state {
	case 0:
		return "hello", nil
	case 1:
		return "world", nil
	default:
		return "", seqreader.ErrEndOfSeq
	}
}
