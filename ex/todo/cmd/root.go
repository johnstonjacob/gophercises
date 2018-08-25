package cmd

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // comment justifying it?
	"github.com/johnstonjacob/gophercises/ex/todo/ink"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Simple CLI Todo List",
	Long:  `Simple CLI Todo List`,
}

var Db, err = gorm.Open("sqlite3", "todo.db")

type Todo struct {
	gorm.Model
	Task     string
	Complete bool
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ink.Error(err.Error())
		os.Exit(1)
	}
	defer Db.Close()
}

func init() {

	if err != nil {
		panic("Failed to connect to DB")
	}

	Db.AutoMigrate(&Todo{})

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(listCmd)
}
