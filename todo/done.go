package todo

import (
	"fmt"
)

// ChangeTaskStatus changes the task status with done[x]/undone[ ]
func ChangeTaskStatus(path string, id int, status string) error {
	// Read from file
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Update status
	id = id - 1
	if id < len(tasks) {
		tasks[id].Done = status
	} else {
		return fmt.Errorf("\nInvalid Task ID - %d \n", id+1)
	}

	// Write to file
	err = WriteTasks(path, tasks)
	if err != nil {
		return err
	}
	if status == "x" {
		fmt.Printf("\nDone: \n    %s \n\n", tasks[id].Name)
	} else {
		fmt.Printf("\nUndone: \n    %s \n\n", tasks[id].Name)
	}

	return nil
}
