package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

var (
	simple     int  // Count for simplified flag
	undoneOnly bool // bool for simplified flag
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list [-s|todo only]",
	Aliases: []string{"l", "ls"},
	Short:   "[l] List all todos.",
	Long:    `[l] List all todos.`,
	Example: `
  todo list
  todo list -s
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if simple >= 1 { // if flag is provided
			undoneOnly = true
		}

		err := todo.ListTasks(todoFile, tag, undoneOnly)
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().CountVarP(&simple, "", "s", "Only todo tasks would be showed")
	listCmd.Flags().StringVarP(&tag, "tag", "t", "", "Search terms to filter result")
}
