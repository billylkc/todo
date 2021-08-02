package todo

import "fmt"

// RemoveTask removes a task by id
func RemoveTask(path string, id int) error {
	var removed Task
	// Read from file
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Update status
	id = id - 1
	if id < len(tasks) {
		removed = tasks[id]
		tasks = append(tasks[:id], tasks[id+1:]...)
	} else {
		return fmt.Errorf("\nInvalid Task ID - %d \n", id+1)
	}

	// Write to file
	err = WriteTasks(path, tasks)
	if err != nil {
		return err
	}
	fmt.Printf("\nRemoved: \n    %s \n\n", removed.Name)

	return nil
}
