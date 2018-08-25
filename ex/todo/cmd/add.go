package cmd

import (
	"strings"

	"github.com/johnstonjacob/gophercises/ex/todo/ink"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a Todo list item",
	Run: func(cmd *cobra.Command, args []string) {
		ink.Success(strings.Join(args, " "))
		ink.Success("Todo list item added")
	},
}
