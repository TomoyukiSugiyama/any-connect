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
	Use:   "config",
	Short: "config",
	Run: func(cmd *cobra.Command, args []string) {
		config()
	},
}

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

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "state",
	Run: func(cmd *cobra.Command, args []string) {
		state()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(disconnectCmd)
	rootCmd.AddCommand(stateCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
