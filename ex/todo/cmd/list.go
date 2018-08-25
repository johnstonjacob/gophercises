package cmd

import (
	"github.com/johnstonjacob/gophercises/ex/todo/ink"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints all Todo list items",
	Run: func(cmd *cobra.Command, args []string) {
		ink.Warning("Todo list items:")
	},
}
