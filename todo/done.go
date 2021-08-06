package todo

import (
	"fmt"
)

// ChangeTaskStatus changes the task status with done[x]/undone[ ]
func ChangeTaskStatus(path string, ids []int, status string) error {
	var jobs []string

	// Read from file
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Create remove map
	statusMap := make(map[int]bool)
	for _, id := range ids {
		id = id - 1 // index 0
		statusMap[id] = true
	}

	// Update status
	for i, task := range tasks {
		if _, ok := statusMap[i]; ok {
			tasks[i].Done = status
			jobs = append(jobs, task.Name)
		}
	}

	// Write to file
	err = WriteTasks(path, tasks, true)
	if err != nil {
		return err
	}

	// Print status
	if status == "x" {
		fmt.Printf("\nDone: \n")
	} else {
		fmt.Printf("\nUndone: \n")
	}
	for _, job := range jobs {
		fmt.Printf("\t%s\n", job)
	}

	return nil
}
