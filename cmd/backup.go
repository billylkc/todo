package cmd

import (
	"github.com/billylkc/todo/todo"
	"github.com/spf13/cobra"
)

var (
	doneCount int
	doneOnly  bool
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"bk"},
	Short:   "[bk] Archive todo tasks",
	Long:    `[bk] Archive todo tasks`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Derive flag
		if doneCount >= 1 {
			doneOnly = true
		}

		// Backup file
		err := todo.BackupTasks(todoFile, doneOnly)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.PersistentFlags().CountVarP(&doneCount, "done", "", "Only done tasks would be archived")
}
