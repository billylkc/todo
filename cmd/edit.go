package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit <id> <todo>",
	Aliases: []string{"e"},
	Short:   "[e] Edit a todo.",
	Long:    `[e] Edit a todo.`,
	Example: `
  todo edit 2 Eat Apple
  todo edit -d 0804 2 Eat Apple
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Handle input date
		var err error
		taskDate, err = todo.ParseInputDate(taskDate)
		if err != nil {
			return err
		}

		var name string
		if len(args) >= 2 {
			name = strings.Join(args[1:], " ") // Second args onwards are updated task
		} else {
			return fmt.Errorf("Invalid number of arguments \n")
		}

		// Check id
		id, err := todo.ConvertID([]string{args[0]}) // First arg is task id
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		// Edit
		err = todo.EditTask(todoFile, id, name, taskDate)
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
	rootCmd.AddCommand(editCmd)
	today := time.Now().Format("2006-01-02")
	editCmd.Flags().StringVarP(&taskDate, "date", "d", today, "Task Date")
}
