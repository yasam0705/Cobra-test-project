package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var countNum int

func init() {
	countByCmd.Flags().IntVarP(&countNum, "count", "c", 1, "integer number")
	sayByeCmd.AddCommand(countByCmd)
	rootCmd.AddCommand(sayByeCmd)
}

var sayByeCmd = &cobra.Command{
	Use:   "saybye [name string]",
	Short: "Say bye",
	Long:  "This command say bye user",
	Args:  cobra.MinimumNArgs(1),
	Run:   sayBye,
}

var countByCmd = &cobra.Command{
	Use:   "count [name string]",
	Short: "Count output text",
	Long:  " ",
	Args:  cobra.MinimumNArgs(1),
	Run:   count,
}

func count(cmd *cobra.Command, args []string) {
	for i := 0; i < countNum; i++ {
		fmt.Println("Bye", strings.Join(args, " "))
	}
}

func sayBye(cmd *cobra.Command, args []string) {
	fmt.Println("Bye", strings.Join(args, " "))
}
