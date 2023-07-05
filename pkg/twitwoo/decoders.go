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
	// Manifest
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.UserInfo",
		"AccountID",
		stringToInt64("decode account id"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ArchiveInfo",
		"SizeBytes",
		stringToInt64("decode size bytes"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ArchiveInfo",
		"MaxPartSizeBytes",
		stringToInt64("decode max part size bytes"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ArchiveInfo",
		"GenerationDate",
		stringToTime("decode generation date", time.RFC3339),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DataFile",
		"Count",
		stringToInt64("decode count"),
	)

	// Tweets
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"FavoriteCount",
		stringToInt64("decode favourite count"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"RetweetCount",
		stringToInt64("decode retweet count"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"CreatedAt",
		stringToTime("decode created at", time.RubyDate),
	)
}

type decoder interface {
	decode(jsoniter.Any)
}

func decode[T decoder](el jsoniter.Any, dest T) {
	dest.decode(el)
}
