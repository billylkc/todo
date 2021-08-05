package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// wipeCmd represents the wipe command
var wipeCmd = &cobra.Command{
	Use:     "wipe",
	Short:   "Wipe all todos",
	Long:    `Wipe all todos`,
	Example: "  todo wipe",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := todo.WipeAll(todoFile)
		if err != nil {
			return err
		}
		fmt.Printf("\nRemoved All Tasks\n\n")

		// List task
		err = todo.ListTasks(todoFile, false)
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(wipeCmd)
}
