package twitwoo

import (
	"io"
	"net"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/afero"
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
	el.ToVal(acip)
}

func (d *Data) readAccountCreationIPs() (io.ReadCloser, int64, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, 0, err
	}

	files := make([]io.Reader, len(m.DataTypes.AccountCreationIP.Files))
	count := int64(0)
	for i, df := range m.DataTypes.AccountCreationIP.Files {
		df := df
		var r afero.File
		r, err = d.readDataFile(&df)
		if err != nil {
			return nil, 0, err
		}
		files[i] = r
		count += df.Count
	}

	return newMultiReadCloser(files...), count, nil
}

// Manifest is the manifest.json file in the archive.
func (d *Data) AccountCreationIPs() ([]AccountCreationIP, error) {
	r, count, err := d.readAccountCreationIPs()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	acips := make([]AccountCreationIP, 0, count)
	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, parseBufSize)
	for iter.ReadArray() {
		var acip AccountCreationIP
		decode(iter.ReadAny().Get("accountCreationIp"), &acip)
		acips = append(acips, acip)
	}

	return acips, nil
}

// EachAccountCreationIP calls fn for each account creation ip.
func (d *Data) EachAccountCreationIP(fn func(AccountCreationIP) error) error {
	r, _, err := d.readAccountCreationIPs()
	if err != nil {
		return err
	}
	defer r.Close()

	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, parseBufSize)
	for iter.ReadArray() {
		var acip AccountCreationIP
		decode(iter.ReadAny().Get("accountCreationIp"), &acip)
		if err = fn(acip); err != nil {
			return err
		}
	}

	return nil
}
