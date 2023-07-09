package twitwoo_test

import (
	"encoding/json"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

func TestEachRaw(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata")
	data := twitwoo.New(fs)
	man, err := data.Manifest()
	require.NoError(t, err)

	items := make([]map[string]any, 0, 5)
	err = twitwoo.EachRaw(data, man.DataTypes.Tweets, func(m map[string]any) error {
		items = append(items, m)
		if len(items) >= 5 {
			return twitwoo.ErrBreak
		}
		return nil
	})
	require.NoError(t, err)
	require.Len(t, items, 5)
	out, err := json.MarshalIndent(items[0], "", "  ")
	require.NoError(t, err)
	require.Contains(t, string(out), `"id": "1629012888185610242"`)
}
