package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the done command
var removeCmd = &cobra.Command{
	Use:     "remove <id>",
	Aliases: []string{"r"},
	Short:   "[r] Remove a todo.",
	Long:    `[r] Remove a todo.`,
	Example: "  todo remove 2",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check id
		ids, err := todo.ConvertIDs(args)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		err = todo.RemoveTask(todoFile, ids)
		if err != nil {
			fmt.Println(err.Error())
		}

		// List task
		err = todo.ListTasks(todoFile, "", false)
		if err != nil {
			fmt.Println(err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
