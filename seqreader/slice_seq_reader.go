package seqreader

// ItemsFromSliceReader adapts a [SeqReader] over slices into a [SeqReader]
// over individual items from those slices, as if all concatenated together.
//
// For example, this can adapt a SeqReader[[]byte] into a SeqReader[byte],
// making the sequence of byte slices appear as a flat sequence of bytes.
//
// The given sequence must either have a finite number of items or must
// eventually produce a non-empty item, or reading from the resulting
// sequence will fail to terminate.
func ItemsFromSliceReader[T any](r SeqReader[[]T]) SeqReader[T] {
	return &itemsFromSliceReader[T]{r, nil}
}

type itemsFromSliceReader[T any] struct {
	r       SeqReader[[]T]
	current []T
}

func (r *itemsFromSliceReader[T]) ReadSeq() (T, error) {
	for len(r.current) == 0 {
		var err error
		r.current, err = r.r.ReadSeq()
		if err != nil {
			var zero T
			return zero, err
		}
	}
	ret, remain := r.current[0], r.current[1:]
	r.current = remain
	return ret, nil
}
