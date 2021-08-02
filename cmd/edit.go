package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit <id> <todo>",
	Aliases: []string{"e"},
	Short:   "[e] Edit a todo.",
	Long:    `[e] Edit a todo.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check id
		id, err := todo.ConvertID(args)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		var name string
		if len(args) == 2 {
			name = args[1]
			fmt.Println(name)
		}

		err = todo.EditTask(todoFile, id, name)
		if err != nil {
			fmt.Println(err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

}
