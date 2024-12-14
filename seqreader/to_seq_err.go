package seqreader

import (
	"iter"
)

// ToSeq adapts the given [SeqReader] into an object that provides both
// an [iter.Seq] and a possible error encountered while reading from it.
//
// This is to allow using a [SeqReader] with a function that expects
// [iter.Seq], by splitting the iteration responsibility from the
// error-handling responsibility.
func ToSeq[T any](r SeqReader[T]) *FallibleSeq[T] {
	return &FallibleSeq[T]{r, nil}
}

// FallibleSeq provides an [iter.Seq] over items of type T where reads
// can potentially fail with an error.
//
// After exhausting the sequence, call method Err to determine whether
// an error occured:
//
//	for item := range fallibleSeq.Items() {
//	    // Do something with item
//	}
//	if err := fallibleSeq.Err(); err != nil {
//	    // Handle err
//	}
type FallibleSeq[T any] struct {
	r   SeqReader[T]
	err error
}

// Items returns an [iter.Seq] over the items in the sequence, which
// terminates either when the underlying sequence is exhausted or when
// reading from that sequence returns an error.
//
// Call method Err after the sequence is exhausted to determine whether
// an error occurred.
func (s *FallibleSeq[T]) Items() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			next, err := s.r.ReadSeq()
			if err == ErrEndOfSeq {
				return
			}
			s.err = err
			if err != nil {
				return
			}
			if !yield(next) {
				return
			}
		}
	}
}

func (s *FallibleSeq[T]) Err() error {
	return s.err
}
