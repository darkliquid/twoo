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
	fi, err := os.Stat(path)
	if err != nil {
		return nil, nil, err
	}

	if fi.IsDir() {
		fs := afero.NewBasePathFs(afero.NewOsFs(), path)
		return fs, func() error { return nil }, nil
	}

	var r *zip.ReadCloser
	r, err = zip.OpenReader(path)
	if err != nil {
		return nil, nil, err
	}

	return zipfs.New(&r.Reader), r.Close, nil
}
