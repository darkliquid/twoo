package twitwoo

import (
	"strconv"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func stringToInt64(op string) jsoniter.DecoderFunc {
	return func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		var err error
		*((*int64)(ptr)), err = strconv.ParseInt(iter.ReadString(), 10, 64)
		if err != nil {
			iter.ReportError(op, err.Error())
			return
		}
	}
}

func stringToTime(op, layout string) jsoniter.DecoderFunc {
	return func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		var err error
		*((*time.Time)(ptr)), err = time.Parse(layout, iter.ReadString())
		if err != nil {
			iter.ReportError(op, err.Error())
			return
		}
	}
}

func init() {
	registerManifestDecoders()
	registerTweetDecoders()
}

type decoder interface {
	decode(jsoniter.Any)
}
