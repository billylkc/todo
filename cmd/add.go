package cmd

import (
	"fmt"
	"strings"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <todo>",
	Aliases: []string{"a"},
	Short:   "[a] Add a todo.",
	Long:    `[a] Add a todo.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) >= 1 {
			task := strings.Join(args, " ")
			err := todo.AddTask(todoFile, task)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("\nAdded task - %s \n\n", task)
		}

		// List task
		err := todo.ListTasks(todoFile)
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
