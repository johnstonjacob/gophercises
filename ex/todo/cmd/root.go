package cmd

import (
	"os"

	"github.com/johnstonjacob/gophercises/ex/todo/ink"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Simple CLI Todo List",
	Long:  `Simple CLI Todo List`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ink.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
