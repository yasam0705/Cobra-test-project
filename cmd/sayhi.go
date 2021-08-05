package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sayHiCmd)
}

var sayHiCmd = &cobra.Command{
	Use:   "sayhi [name string]",
	Short: "Say hi",
	Long:  "This command say hi user",
	Args:  cobra.MinimumNArgs(1),
	Run:   sayHi,
}

func sayHi(cmd *cobra.Command, args []string) {
	fmt.Println("Hi", strings.Join(args, " "))
}
