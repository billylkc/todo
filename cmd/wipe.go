package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// wipeCmd represents the wipe command
var wipeCmd = &cobra.Command{
	Use:   "wipe",
	Short: "Wipe all todos",
	Long:  `Wipe all todos`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := todo.WipeAll(todoFile)
		if err != nil {
			return err
		}
		fmt.Printf("\nRemoved All Tasks\n\n")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(wipeCmd)
}
