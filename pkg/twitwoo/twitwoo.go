package twitwoo

import (
	"fmt"

	"github.com/spf13/afero"
)

// Data is a struct that provides an interface to the data archive.
type Data struct {
	fs           afero.Fs
	manifestData *Manifest
}

// New returns a new Data struct.
func New(fs afero.Fs) *Data {
	return &Data{
		fs: fs,
	}
}

func (d *Data) readDataFile(df *DataFile) (afero.File, error) {
	f, err := d.fs.Open(df.Name)
	if err != nil {
		return nil, err
	}

	if err = skipPreamble(fmt.Sprintf("window.%s = ", df.Preamble), f); err != nil {
		return nil, err
	}

	return f, nil
}
