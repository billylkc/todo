package todo

import (
	"fmt"
)

// BackupTasks reads from the original task file and append the content to the backup file
// If doneOnly flag is provided, only the finished tasks will be back up
func BackupTasks(path string, doneOnly bool) error {
	var allTask []Task       // task from the original file
	var originalTasks []Task // original task that would be written back to the original files / undone tasks
	var bkTasks []Task       // done tasks to be backuped / done tasks

	// Read tasks
	allTask, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Handle doneOnly flag and backup all
	for _, task := range allTask {
		if doneOnly {
			if task.Done == "x" {
				bkTasks = append(bkTasks, task)
			} else {
				originalTasks = append(originalTasks, task)
			}
		} else {
			// Backup all
			bkTasks = append(bkTasks, task)
		}
	}

	if len(bkTasks) == 0 {
		return fmt.Errorf("/nNo new todo items for archive. /n")
	}

	// Append to backup files
	backupPath := path + ".bk"                   // ~/todo.txt.bk
	err = WriteTasks(backupPath, bkTasks, false) // append
	if err != nil {
		return fmt.Errorf("Can not write to backup file - %s \n", backupPath)
	}

	// Write back to oringal file if any (undone tasks)
	if len(originalTasks) > 0 {
		err = WriteTasks(path, originalTasks, true) // overwrite
		if err != nil {
			return fmt.Errorf("Can not write back to todo file - %s \n", backupPath)
		}

		fmt.Printf("\nWritten %d lines to todo file [%s]\n\n", len(originalTasks), path)
	}
	fmt.Printf("\nWritten %d lines to backup file [%s]\n\n", len(bkTasks), backupPath)

	return nil
}
