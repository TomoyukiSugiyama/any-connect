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

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect",
	Run: func(cmd *cobra.Command, args []string) {
		connect()
	},
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "disconnect",
	Run: func(cmd *cobra.Command, args []string) {
		disconnect()
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(disconnectCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
