package todo

import "fmt"

// EditTask edit the corresponsing task to a new one
// Same complete status as the original task
func EditTask(path string, id int, name string) error {

	// Read old tasks and add new ones
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Update Task
	var oldtask, newtask Task
	id = id - 1
	if id < len(tasks) {
		for i, _ := range tasks {
			if i == id {
				oldtask = tasks[i]
				newtask = Task{
					Date: oldtask.Date,
					Name: name,
					Done: oldtask.Done,
				}
				tasks[i] = newtask
			}
		}

	} else {
		return fmt.Errorf("\nInvalid Task ID - %d \n", id+1)
	}

	// Write to file
	err = WriteTasks(path, tasks)
	if err != nil {
		return err
	}

	return nil
}
