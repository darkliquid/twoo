package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerReplyPromptDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ReplyPrompt",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ReplyPrompt is the structure of the data/reply-prompt.js file.
type ReplyPrompt struct {
    // Fields go here
}

func (r *ReplyPrompt) decode(el jsoniter.Any) {
	el = el.Get("replyPrompt")
	el.ToVal(r)
}

// ReplyPrompts returns all the ReplyPrompt items.
func (d *Data) ReplyPrompts() ([]*ReplyPrompt, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ReplyPrompt](d, m.DataTypes.ReplyPrompt)
}

// EachReplyPrompt calls fn for each ReplyPrompt item.
func (d *Data) EachReplyPrompt(fn func(*ReplyPrompt) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ReplyPrompt](d, m.DataTypes.ReplyPrompt, fn)
}
