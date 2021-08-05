package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <todo>",
	Aliases: []string{"a"},
	Short:   "[a] Add a todo.",
	Long:    `[a] Add a todo.`,
	Example: `
  todo add Eat Apple
  todo add -d 0804 Eat Apple
  `,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Handle input date
		var err error
		taskDate, err = todo.ParseInputDate(taskDate)
		if err != nil {
			return err
		}

		if len(args) >= 1 {
			task := strings.Join(args, " ")
			err := todo.AddTask(todoFile, task, taskDate)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("\nAdded task - %s \n", task)
		} else {
			return fmt.Errorf("\n\nPlease add at least one task.\n")
		}

		// List task
		err = todo.ListTasks(todoFile, "", false) // all
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	today := time.Now().Format("2006-01-02")
	addCmd.Flags().StringVarP(&taskDate, "date", "d", today, "Task date")
}
