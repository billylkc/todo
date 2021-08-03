package todo

import (
	"fmt"
)

// EditTask edits the corresponsing task to a new one
func EditTask(path string, id int, name string, date string) error {

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
					Date: date,
					Name: name,
					Done: " ",
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

	// Print
	msg := `
Edited:
  %s

To be:
  %s

`
	fmt.Printf(msg, oldtask.Name, newtask.Name)

	return nil
}
