package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(factorialCmd)
}

var factorialCmd = &cobra.Command{
	Use:   "fact [int number]",
	Short: "Calculte factorial",
	Long:  "Calculate factorial with for",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		n, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var result int = 1
		for i := 1; i <= n; i++ {
			result *= i
		}
		fmt.Printf("Factorial %d: %d\n", n, result)
	},
}
