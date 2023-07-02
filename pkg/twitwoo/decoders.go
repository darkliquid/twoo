package twitwoo

import (
	"strconv"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterFieldDecoderFunc("twitwoo.Tweet", "FavoriteCount", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		var err error
		*((*int64)(ptr)), err = strconv.ParseInt(iter.ReadString(), 10, 64)
		if err != nil {
			iter.ReportError("decode favorite count", err.Error())
			return
		}
	})

	jsoniter.RegisterFieldDecoderFunc("twitwoo.Tweet", "RetweetCount", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		var err error
		*((*int64)(ptr)), err = strconv.ParseInt(iter.ReadString(), 10, 64)
		if err != nil {
			iter.ReportError("decode retweet count", err.Error())
			return
		}
	})

	jsoniter.RegisterFieldDecoderFunc("twitwoo.Tweet", "CreatedAt", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		timeStr := iter.ReadString()
		t, err := time.Parse(time.RubyDate, timeStr)
		if err != nil {
			iter.ReportError("decode created at", err.Error())
			return
		}

		*((*time.Time)(ptr)) = t
	})
}

type decoder interface {
	decode(jsoniter.Any)
}

func decode[T decoder](el jsoniter.Any, dest T) {
	dest.decode(el)
}
