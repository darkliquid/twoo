package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SubcommandExactArgs returns an error if there are not exactly n args
// after one of the given list of subcommands.
func SubcommandExactArgs(commands []string, n int) cobra.PositionalArgs {
	return func(_ *cobra.Command, args []string) error {
		if len(args) != n+1 {
			return fmt.Errorf("accepts %d arg(s), received %d", n+1, len(args))
		}

		if len(args) > 0 {
			for _, arg := range commands {
				if args[0] == arg {
					return nil
				}
			}
			return fmt.Errorf("subcommand %s not recognised", args[0])
		}
		return nil
	}
}
