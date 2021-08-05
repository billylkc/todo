package cmd

import (
	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

var tag string // search terms for summary

// summarizeCmd represents the summerize command
var summarizeCmd = &cobra.Command{
	Use:     "summarize",
	Aliases: []string{"s"},
	Short:   "[s] Summarize man-day of todo tasks",
	Long:    `[s] Summarize man-day of todo tasks`,
	Example: `
  todo summarize
  todo summarize -t Meeting
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := todo.SummarizeTask(todoFile, tag)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
	summarizeCmd.Flags().StringVarP(&tag, "tag", "t", "", "Search terms to filter result")
}
