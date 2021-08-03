package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// undoneCmd represents the done command
var undoneCmd = &cobra.Command{
	Use:     "undone <id>",
	Aliases: []string{"u"},
	Short:   "[u] Unone a todo.",
	Long:    `[u] Undone a todo.`,
	Example: "  todo undone 2",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check id
		id, err := todo.ConvertID(args)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		err = todo.ChangeTaskStatus(todoFile, id, " ")
		if err != nil {
			fmt.Println(err.Error())
		}

		// List task
		err = todo.ListTasks(todoFile)
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
