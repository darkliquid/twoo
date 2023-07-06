package twitwoo

import (
	"net"

	jsoniter "github.com/json-iterator/go"
)

func registerAccountCreationIPDecoders() {
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AccountCreationIP",
		"AccountID",
		stringToInt64("decode account id"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AccountCreationIP",
		"IP",
		stringToIP("decode IP"),
	)
}

// AccountCreationIP is the structure of the account-creation-ip.json file.
type AccountCreationIP struct {
	IP        net.IP `json:"userCreationIp"`
	AccountID int64  `json:"accountId"`
}

func (acip *AccountCreationIP) decode(el jsoniter.Any) {
	el = el.Get("accountCreationIp")
	el.ToVal(acip)
}

// AccountCreationIPs returns all the account creation ips.
func (d *Data) AccountCreationIPs() ([]*AccountCreationIP, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AccountCreationIP](d, m.DataTypes.AccountCreationIP)
}

// EachAccountCreationIP calls fn for each account creation ip.
func (d *Data) EachAccountCreationIP(fn func(*AccountCreationIP) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AccountCreationIP](d, m.DataTypes.AccountCreationIP, fn)
}
