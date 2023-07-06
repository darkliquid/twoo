package extract

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var acipFormat string

const defaultAcipFormat = "{{.IP}}\n"

// acipCmd represents the extract command.
var acipCmd = &cobra.Command{
	Use:   "account-creation-ips FILE|DIR",
	Short: "extract account creation ips",
	Long:  `extract the account creation ips included in the archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		fs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		tmpl, err := template.New("acip").Parse(acipFormat)
		if err != nil {
			return err
		}

		data := twitwoo.New(fs)
		return data.EachAccountCreationIP(func(acip twitwoo.AccountCreationIP) error {
			return tmpl.Execute(os.Stdout, acip)
		})
	},
}

func init() {
	acipCmd.Flags().StringVarP(&acipFormat, "format", "f", defaultAcipFormat, "format of extracted account creation ips")
}
