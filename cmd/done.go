package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done <id>",
	Aliases: []string{"d"},
	Short:   "[d] Done a todo.",
	Long:    `[d] Done a todo,`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check id
		id, err := todo.ConvertID(args)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		// Change to done
		err = todo.ChangeTaskStatus(todoFile, id, "x")
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
	rootCmd.AddCommand(doneCmd)
}
