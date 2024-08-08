package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "any-connect [sub]",
		Short: "Any Connect CLI",
	}
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create or overwrite a config file ( ~/.cisco_vpn/credentials )",
	Run: func(cmd *cobra.Command, args []string) {
		config()
	},
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a virtual private network",
	Run: func(cmd *cobra.Command, args []string) {
		connect()
	},
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from a virtual private network",
	Run: func(cmd *cobra.Command, args []string) {
		disconnect()
	},
}

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Check state ( connected or disconnected )",
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
