package seqreader

import (
	"errors"
	"io"
)

// SeqReader is a sequence of items of type T from which reads can potentially
// fail.
//
//	for {
//	    item, err := seqReader.ReadSeq()
//	    if err == ErrEndOfSeq {
//	        break
//	    }
//	    if err != nil {
//	        // handle err
//	    }
//	    // handle item
//	}
type SeqReader[T any] interface {
	// ReadSeq either returns the next item in the sequence or returns an
	// error describing why another item cannot be read.
	//
	// At the end of the sequence the error is [ErrEndOfSeq].
	ReadSeq() (T, error)
}

// SeqReadCloser is a [SeqReader] that also implements [io.Close], which
// the caller should call once they are finished reading items from the
// sequence.
type SeqReadCloser[T any] interface {
	SeqReader[T]
	io.Closer
}

// ErrEndOfSeq is the error returned by [SeqReader.ReadSeq] once the end of
// a finite sequence has been reached.
var ErrEndOfSeq = errors.New("end of sequence")
