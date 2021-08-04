package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "Test CLI project",
	Long:  "My first CLI project using cobra",
}
