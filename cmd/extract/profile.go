package extract

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var profileFormat string

const defaultProfileFormat = `Bio: {{.Description.Bio}}
Website: {{.Description.Website}}
Location: {{.Description.Location}}
`

// profileCmd represents the extract command.
var profileCmd = &cobra.Command{
	Use:   "profile FILE|DIR",
	Short: "extract account profile",
	Long:  `extract the account profile included in the archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		fs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		tmpl, err := template.New("profile").Parse(profileFormat)
		if err != nil {
			return err
		}

		data := twitwoo.New(fs)
		return data.EachProfile(func(profile *twitwoo.Profile) error {
			return tmpl.Execute(os.Stdout, profile)
		})
	},
}

func init() {
	profileCmd.Flags().StringVarP(&profileFormat, "format", "f", defaultProfileFormat, "format of extracted account profile")
}
