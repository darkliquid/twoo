#!/bin/bash

datatypes=$(grep 'DataType `json:' manifest.go | awk '{ print $1 }')

for datatype in $datatypes; do
    snakecase=$(echo $datatype | sed 's/\([^A-Z]\)\([A-Z0-9]\)/\1_\2/g' | sed 's/\([A-Z0-9]\)\([A-Z0-9]\)\([^A-Z]\)/\1_\2\3/g' | tr '[:upper:]' '[:lower:]')
    gofile="$snakecase.go"

    datafile=$(echo $snakecase | sed 's/_/-/g')
    datafile="data/$datafile.js"

    datafield=$(echo $datatype | sed 's/\(^[A-Z]\)/\L\1/g' | sed -E 's/([A-Z])([A-Z]+)/\1\L\2/g')

    if [ ! -f $gofile ]; then
        cat >$gofile <<EOF
package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func register${datatype}Decoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.${datatype}",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// $datatype is the structure of the $datafile file.
type $datatype struct {
    // Fields go here
}

func (${datafield:0:1} *$datatype) decode(el jsoniter.Any) {
	el = el.Get("$datafield")
	el.ToVal(${datafield:0:1})
}

// ${datatype}s returns all the $datatype items.
func (d *Data) ${datatype}s() ([]*$datatype, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*$datatype](d, m.DataTypes.$datatype)
}

// Each$datatype calls fn for each $datatype item.
func (d *Data) Each$datatype(fn func(*$datatype) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*$datatype](d, m.DataTypes.$datatype, fn)
}
EOF
    fi
done
