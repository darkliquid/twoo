package twitwoo

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/afero"
)

type Ptr[T any] interface {
	*T
	decoder
}

func readDataType(d *Data, dt DataType) (io.ReadCloser, int64, error) {
	files := make([]io.Reader, len(dt.Files))
	count := int64(0)
	for i, df := range dt.Files {
		df := df
		var r afero.File
		r, err := d.readDataFile(&df)
		if err != nil {
			return nil, 0, err
		}
		files[i] = r
		count += df.Count
	}

	return newMultiReadCloser(files...), count, nil
}

// All returns all items of a data type.
func All[T Ptr[U], U any](d *Data, dt DataType) ([]T, error) {
	r, count, err := readDataType(d, dt)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	items := make([]T, 0, count)
	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, parseBufSize)
	for iter.ReadArray() {
		item := T(new(U))
		item.decode(iter.ReadAny())
		items = append(items, item)
	}

	return items, nil
}

// Each calls fn for each data type item.
func Each[T Ptr[U], U any](d *Data, dt DataType, fn func(T) error) error {
	r, _, err := readDataType(d, dt)
	if err != nil {
		return err
	}
	defer r.Close()

	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, parseBufSize)
	for iter.ReadArray() {
		item := T(new(U))
		item.decode(iter.ReadAny())
		if err = fn(item); err != nil {
			return err
		}
	}

	return nil
}
