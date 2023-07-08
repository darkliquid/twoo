package fs

import (
	"io/fs"

	"github.com/spf13/afero"
)

type aferoFS struct {
	afero.Fs
}

func (fs aferoFS) Open(name string) (fs.File, error) {
	return fs.Fs.Open(name)
}

func AferoFS(fs afero.Fs) fs.FS {
	return aferoFS{fs}
}
