package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	date    = "2/6/2025"
)

// https://stackoverflow.com/questions/4842424/list-of-ansi-color-escape-sequences
// https://patorjk.com/software/taag/#p=display&f=Ivrit&t=Zero

//go:embed logo.txt
var projASCIILogo string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show zero version and build date",
	RunE: func(_ *cobra.Command, _ []string) error {
		fmt.Println(projASCIILogo)
		fmt.Printf(" Version: %s\n", version)
		fmt.Printf("    Date: %s\n", date)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
