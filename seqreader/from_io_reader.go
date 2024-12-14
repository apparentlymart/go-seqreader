package seqreader

import (
	"io"
)

// FromIOReader returns a [SeqReader] over byte slices returned from
// the given reader.
//
// The resulting reader needs byte arrays to recieve data from the [io.Reader].
// makeBuf should return a slice whose length is the maximum size of a single
// read but which may have excess capacity to allow using the remainder of
// the underlying array for future reads as long as the remaining capacity is
// at least the original slice length.
//
// The [SeqReader] returns [ErrEndOfSeq] when the underlying [io.Reader]
// returns [io.EOF].
//
// The result dynamically implmenets [SeqReadCloser], calling Close on
// the given reader if it is actually an [io.ReadCloser], or a no-op success
// if not.
func FromIOReader(r io.Reader, makeBuf func() []byte) SeqReader[[]byte] {
	initialBuf, maxRead := ioReaderBuf(makeBuf)
	return &ioReaderSeq{
		r:       r,
		makeBuf: makeBuf,
		buf:     initialBuf,
		maxRead: maxRead,
	}
}

type ioReaderSeq struct {
	r       io.Reader
	makeBuf func() []byte
	buf     []byte
	maxRead int
}

func (s *ioReaderSeq) ReadSeq() ([]byte, error) {
	if len(s.buf) < s.maxRead {
		s.buf, s.maxRead = ioReaderBuf(s.makeBuf)
	}
	n, err := s.r.Read(s.buf[:s.maxRead])
	ret := s.buf[:n]
	s.buf = s.buf[n:]
	if err == io.EOF {
		err = ErrEndOfSeq
	}
	return ret, err
}

func (s *ioReaderSeq) Close() error {
	closer, ok := s.r.(io.Closer)
	if !ok {
		return nil
	}
	return closer.Close()
}

func ioReaderBuf(makeBuf func() []byte) (buf []byte, maxRead int) {
	buf = makeBuf()
	maxRead = len(buf)
	buf = buf[:cap(buf)]
	return buf, maxRead
}
