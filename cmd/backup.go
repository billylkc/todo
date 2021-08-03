package cmd

import (
	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"bk"},
	Short:   "[bk] Archive todo tasks",
	Long:    `[bk] Archive todo tasks`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Backup file
		err := todo.BackupTasks(todoFile)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
