package twitwoo

import (
	"errors"
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

// ErrBreak exits the Each iterator.
var ErrBreak = errors.New("end each iterator")

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
			if err == ErrBreak {
				return nil
			}
			return err
		}
	}

	return nil
}

// EachRaw calls fn for each data type item, but all items are map[string]any.
func EachRaw(d *Data, dt DataType, fn func(map[string]any) error) error {
	r, _, err := readDataType(d, dt)
	if err != nil {
		return err
	}
	defer r.Close()

	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, parseBufSize)
	for iter.ReadArray() {
		item := make(map[string]any)

		// Every item is an object with a single key of the object type
		// with a value of the _actual_ data type item.
		if iter.ReadObject() != "" {
			iter.ReadAny().ToVal(&item)
			if err = fn(item); err != nil {
				if err == ErrBreak {
					return nil
				}
				return err
			}
		}

		// Need to finish reading the object so the next ReadArray
		// will work.
		if key := iter.ReadObject(); key != "" {
			panic(key)
		}
	}

	return nil
}
