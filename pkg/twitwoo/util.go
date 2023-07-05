package twitwoo

import (
	"bytes"
	"fmt"
	"io"
)

func skipPreamble(preamble string, r io.Reader) error {
	buf := make([]byte, len(preamble))
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}

	if !bytes.Equal(buf, []byte(preamble)) {
		return fmt.Errorf("preamble does not match")
	}

	return nil
}

type multiReadCloser struct {
	io.Reader
	closers []io.Closer
}

// NewMultiReadCloser returns a new io.ReadCloser that's the
// logical concatenation of the provided input readers.
func NewMultiReadCloser(r ...io.Reader) *multiReadCloser {
	mrc := &multiReadCloser{
		Reader: io.MultiReader(r...),
	}

	for _, rc := range r {
		if closer, ok := rc.(io.Closer); ok {
			mrc.closers = append(mrc.closers, closer)
		}
	}

	return mrc
}

// Close closes all the readers. Implements io.Closer.
func (mrc *multiReadCloser) Close() error {
	var err error
	for _, c := range mrc.closers {
		if cerr := c.Close(); cerr != nil {
			err = cerr
		}
	}
	return err
}
