package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var divideCmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"div"},
	Short:   "Divide one number by another",
	Long:    "Carry out division operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := Divide(args[0], args[1], shouldRoundUp)
		if err != nil {
			return err
		}
		fmt.Printf("Division of %s by %s = %s.\n\n", args[0], args[1], res)
		return nil
	},
}

func Divide(divide string, by string, shouldRoundUp bool) (result string, e error) {
	num1, err := strconv.ParseFloat(divide, 64)
	if err != nil {
		return "", fmt.Errorf("first value is not a number")
	}
	num2, err := strconv.ParseFloat(by, 64)
	if err != nil {
		return "", fmt.Errorf("second value is not a number")
	}
	if num2 == 0 {
		return "", fmt.Errorf("cannot divide by zero")
	}
	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1/num2), nil
	}
	return fmt.Sprintf("%f", num1/num2), nil
}

func init() {
	divideCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
	rootCmd.AddCommand(divideCmd)
}
