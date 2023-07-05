package twitwoo

import (
	"bytes"
	"fmt"
	"io"
)

const (
	manifestPreamble = `window.__THAR_CONFIG = `
	tweetsPreamble   = "window.YTD.tweets.part0 = "
)

func SkipPreamble(preamble string, r io.Reader) error {
	buf := make([]byte, len(preamble))
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}

	if !bytes.Equal(buf, []byte(preamble)) {
		return fmt.Errorf("preamble does not match")
	}

	return nil
}
