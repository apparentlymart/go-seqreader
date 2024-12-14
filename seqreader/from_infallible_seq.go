package seqreader

import (
	"iter"
)

// FromSeq adapts the given seq into a [SeqReadCloser].
//
// Calling Close on the result tells the sequence to terminate.
func FromSeq[T any](seq iter.Seq[T]) SeqReadCloser[T] {
	next, stop := iter.Pull(seq)
	return fromSeq[T]{next, stop}
}

type fromSeq[T any] struct {
	next func() (T, bool)
	stop func()
}

func (s fromSeq[T]) ReadSeq() (T, error) {
	ret, ok := s.next()
	if !ok {
		return ret, ErrEndOfSeq
	}
	return ret, nil
}

func (s fromSeq[T]) Close() error {
	s.stop()
	return nil
}
