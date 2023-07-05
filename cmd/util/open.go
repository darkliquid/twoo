package util

import (
	"archive/zip"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/afero/zipfs"
)

// Open opens a file or directory and returns an afero.Fs, a function to
// close the file or directory, and an error if one occurred.
func Open(path string) (afero.Fs, func() error, error) {
	close := func() error {
		return nil
	}

	fi, err := os.Stat(path)
	if err != nil {
		return nil, nil, err
	}

	var fs afero.Fs
	if fi.IsDir() {
		fs = afero.NewBasePathFs(afero.NewOsFs(), path)
	} else {
		r, err := zip.OpenReader(path)
		if err != nil {
			return nil, nil, err
		}
		close = r.Close

		fs = zipfs.New(&r.Reader)
	}

	return fs, close, nil
}
