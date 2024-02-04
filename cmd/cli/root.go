package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var noKeep bool

var rootCmd = &cobra.Command{
	Use:     "allyignore",
	Short:   "A tool that finds and removes unnecessary lines from .gitignore files.",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		Start(noKeep)
	},
}

func Execute() {
	rootCmd.Flags().BoolP("help", "h", false, "Help for Allyignore.")
	rootCmd.Flags().BoolP("version", "v", false, "Version for Allyignore.")
	rootCmd.SetVersionTemplate("Allyignore - Version {{.Version}}\n")
	rootCmd.Flags().BoolVarP(&noKeep, "nokeep", "n", false, "Don't keep comments and empty lines.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
