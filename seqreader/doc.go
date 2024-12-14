// Package seqreader is a utility library for representing fallible sequences.
//
// A fallible sequence is a series of values for which an error could
// potentially occur on any read.
//
// [SeqReader] is therefore a fallible equivalent of [iter.Seq], whose design
// is inspired by [io.Reader]. [iter.Seq] and [iter.Seq2] are both designed for
// infallible sequences and the common patterns using them are not well-suited
// to fallible sequences.
//
// This package also includes various adapter functions to help with using
// [SeqReader] values where [iter.Seq] or [iter.Seq2] values are expected, using
// [iter.Seq] values where [SeqReader] values are expected, and using
// [io.Reader] values where [SeqReader] values are expected.
package seqreader

import (
	_ "io"
	_ "iter"
)
