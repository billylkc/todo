package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list [-a]",
	Aliases: []string{"l"},
	Short:   "[l] List undone todos. If -a given, list all todo.",
	Long:    `[l] List undone todos. If -a given, list all todo.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := todo.ListTasks(todoFile)
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
