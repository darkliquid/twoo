package twitwoo

import (
	"bufio"
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

// dataFile wraps a file and a bufio.Reader and skips past the JS preamble in the file.
type dataFile struct {
	f afero.File
	*bufio.Reader
}

func (df *dataFile) Close() error {
	return df.f.Close()
}

func (d *Data) readData(fn string) (*dataFile, error) {
	f, err := d.fs.Open(fmt.Sprintf("data/%s.js", fn))
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		if b == '=' {
			break
		}
	}

	return &dataFile{f: f, Reader: r}, nil
}
