package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/darkliquid/twoo/cmd/extract"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "twoo",
	Short: "TWitter Offline Online",
	Long:  `An application to take a twitter data archive and make it hostable`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.twoo.yaml)")
	rootCmd.AddCommand(extract.Command())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".twoo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".twoo")
	}

	viper.SetEnvPrefix("TWOO")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match.

	bindArgs("", rootCmd, viper.GetViper())

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func bindArgs(prefix string, cmd *cobra.Command, v *viper.Viper) {
	// Apply config for all flags for the command.
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Build the config file name for this flag setting.
		configName := fmt.Sprintf("%s.%s", prefix, f.Name)

		// Apply the viper config value to the flag if it isn't overridden.
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	// Ensure our prefix has a trailing dot if set.
	if prefix != "" {
		prefix += "."
	}

	// Recursively apply config for all subcommands.
	for _, c := range cmd.Commands() {
		bindArgs(prefix+c.Name(), c, v)
	}
}
