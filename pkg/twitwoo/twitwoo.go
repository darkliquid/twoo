package twitwoo

import (
	"fmt"

	"github.com/spf13/afero"
)

// Data is a struct that provides an interface to the data archive.
type Data struct {
	fs afero.Fs
}

// New returns a new Data struct.
func New(fs afero.Fs) *Data {
	return &Data{
		fs: fs,
	}
}

func (d *Data) readData(fn, preamble string) (afero.File, error) {
	f, err := d.fs.Open(fmt.Sprintf("data/%s.js", fn))
	if err != nil {
		return nil, err
	}

	if err = SkipPreamble(preamble, f); err != nil {
		return nil, err
	}

	return f, nil
}
