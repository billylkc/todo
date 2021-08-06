package todo

import "fmt"

// WipeTasks removes all the tasks
// If doneOnly flag is provided, only the finished tasks will be wiped
func WipeTasks(path string, doneOnly bool) error {
	var (
		tasks []Task
		count int
	)
	if doneOnly {
		allTasks, err := ReadTasks(path)
		if err != nil {
			return err
		}
		for _, task := range allTasks {
			if task.Done == "x" {
				continue
			} else {
				// write back undone task to todd
				tasks = append(tasks, task)
			}
			err := WriteTasks(path, tasks, true)
			if err != nil {
				return err
			}
		}
		count = len(allTasks) - len(tasks)

	} else {
		// wipe all
		err := WriteTasks(path, tasks, true)
		if err != nil {
			return err
		}
	}

	// Print message
	if doneOnly {
		fmt.Printf("\n%d todo tasks removed. \n", count)
	} else {
		fmt.Println("\nAll todo tasks removed.")
	}

	return nil
}
