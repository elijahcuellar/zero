package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mul", "multiple", "multi"},
	Short:   "Multiply 2 numbers",
	Long:    "Carry out multiplication operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Multiplication of %s and %s = %s.\n\n", args[0], args[1], Multiply(args[0], args[1], shouldRoundUp))
	},
}

func Multiply(first string, second string, shouldRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is not a decimal")
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is not a decimal")
		return
	}
	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2)
	}
	return fmt.Sprintf("%f", num1*num2)
}

func init() {
	multiplyCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
	rootCmd.AddCommand(multiplyCmd)
}
