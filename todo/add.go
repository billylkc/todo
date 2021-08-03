package todo

import (
	"encoding/json"
	"strings"
)

// AddTask adds the tasks to the file
func AddTask(path string, name string, date string) error {

	// Read old tasks and add new ones
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	task := Task{
		Date: date,
		Name: strings.ReplaceAll(name, ",", ";"),
		Done: " ",
	}
	tasks = append(tasks, task)

	// Write to file
	err = WriteTasks(path, tasks)
	if err != nil {
		return err
	}

	return nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
