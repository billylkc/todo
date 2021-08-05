package todo

import "fmt"

// RemoveTask removes a task by ids
func RemoveTask(path string, ids []int) error {
	var old []Task
	var new []Task

	// Read from file
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Create remove map
	removeMap := make(map[int]bool)
	for _, id := range ids {
		id = id - 1 // index 0
		removeMap[id] = true
	}

	// Update status
	for i, task := range tasks {
		if _, ok := removeMap[i]; ok {
			old = append(old, task)
		} else {
			new = append(new, task)
		}
	}

	// Write to file
	err = WriteTasks(path, new)
	if err != nil {
		return err
	}
	fmt.Printf("\nRemoved: \n")
	for _, task := range old {
		fmt.Printf("   %s\n", task.Name)
	}

	return nil
}
