package cmd

import (
	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// summarizeCmd represents the summerize command
var summarizeCmd = &cobra.Command{
	Use:     "summarize",
	Aliases: []string{"s"},
	Short:   "[s] Summarize man-day of todo tasks",
	Long:    `[s] Summarize man-day of todo tasks`,
	Example: "  todo summarize",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := todo.SummarizeTask(todoFile)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
}
