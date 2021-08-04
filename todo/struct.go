package todo

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Task is the task specified in the todo file
type Task struct {
	Date string
	Name string
	Done string // x as Done, empty as not done
}

// WriteTasks writes the struct to the file
func WriteTasks(path string, tasks []Task) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil
	}

	// Sort tasks first
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Date < tasks[j].Date
	})

	// Write content
	datawriter := bufio.NewWriter(f)
	if len(tasks) > 0 {
		// normal case, write tasks
		for _, task := range tasks {
			line := fmt.Sprintf("%s,%s,%s", task.Date, task.Name, task.Done)
			_, _ = datawriter.WriteString(line + "\n")
		}
	} else {
		// wipe
		_, _ = datawriter.WriteString("")
	}
	datawriter.Flush()

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
