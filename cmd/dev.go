package cmd

import (
	"fmt"

	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"x"},
	Short:   "[x] Development",
	Long:    `[x] Development`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := todo.Dev()
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
