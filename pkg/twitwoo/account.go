package twitwoo

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

func registerAccountDecoders() {
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Account",
		"AccountID",
		stringToInt64("decode account id"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Account",
		"CreatedAt",
		stringToTime("decode created at", time.RFC3339),
	)
}

// Account is the structure of the data/account.js file.
type Account struct {
	CreatedAt   time.Time `json:"createdAt"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"accountDisplayName"`
	CreatedVia  string    `json:"createdVia"`
	AccountID   int64     `json:"accountId"`
}

func (a *Account) decode(el jsoniter.Any) {
	el = el.Get("account")
	el.ToVal(a)
}

// Accounts returns all the Account items.
func (d *Data) Accounts() ([]*Account, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Account](d, m.DataTypes.Account)
}

// EachAccount calls fn for each Account item.
func (d *Data) EachAccount(fn func(*Account) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Account](d, m.DataTypes.Account, fn)
}
