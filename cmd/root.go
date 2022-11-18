package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "any-connect [sub]",
		Short: "Any Connect CLI",
	}
)

var configCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect",
	Run: func(cmd *cobra.Command, args []string) {
		connect()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func initConfig() {

}

func Execute() error {
	return rootCmd.Execute()
}
