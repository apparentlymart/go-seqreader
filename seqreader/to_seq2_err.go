package seqreader

import (
	"iter"
)

// ToSeq2 returns an [iter.Seq2] over results from the given [SeqReader],
// where each item in the sequence includes both a result and an error.
//
// Use this to adapt a [SeqReader] to a situation that requires an [iter.Seq2].
// For example:
//
//	for item, err := range seqreader.ToSeq2[seqReader] {
//	    if err != nil {
//	        // handle err
//	    }
//	    // handle item
//	}
func ToSeq2[T any](r SeqReader[T]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			next, err := r.ReadSeq()
			if err == ErrEndOfSeq {
				return
			}
			if !yield(next, err) {
				return
			}
		}
	}
}
