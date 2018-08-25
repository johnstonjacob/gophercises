package cmd

import (
	"strings"

	"github.com/johnstonjacob/gophercises/ex/todo/ink"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes an item on your Todo list",
	Run: func(cmd *cobra.Command, args []string) {
		ink.Success(strings.Join(args, " "))
		ink.Success("Todo list item finished")
	},
}
