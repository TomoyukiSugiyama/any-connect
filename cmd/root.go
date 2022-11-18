package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "any-connect [sub]",
	Short: "Any Connect CLI",
}

var initCmd = &cobra.Command{
	Use:   "config",
	Short: "configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initCmd")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initConfig() {

}

func Execute() error {
	return rootCmd.Execute()
}
