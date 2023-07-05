package extract

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var manifestFormat string

const defaultManifestFormat = `{{.}}`

// manifestCmd represents the manifest command.
var manifestCmd = &cobra.Command{
	Use:   "manifest FILE|DIR",
	Short: "extract archive manifest data",
	Long:  `extract the manifest data included in the archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		fs, close, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer close()

		tmpl, err := template.New("manifest").Parse(manifestFormat)
		if err != nil {
			return err
		}

		data := twitwoo.New(fs)
		m, err := data.Manifest()
		if err != nil {
			return err
		}

		return tmpl.Execute(os.Stdout, m)
	},
}

func init() {
	manifestCmd.Flags().StringVarP(&manifestFormat, "format", "f", defaultManifestFormat, "format of extracted manifest")
}
