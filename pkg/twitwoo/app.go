package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAppDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.App",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// App is the structure of the data/app.js file.
type App struct {
    // Fields go here
}

func (a *App) decode(el jsoniter.Any) {
	el = el.Get("app")
	el.ToVal(a)
}

// Apps returns all the App items.
func (d *Data) Apps() ([]*App, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*App](d, m.DataTypes.App)
}

// EachApp calls fn for each App item.
func (d *Data) EachApp(fn func(*App) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*App](d, m.DataTypes.App, fn)
}
