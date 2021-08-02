package todo

import (
	"encoding/json"
	"strings"
	"time"
)

// AddTask adds the tasks to the file
func AddTask(path string, name string) error {

	// Read old tasks and add new ones
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}
	today := time.Now().Format("2006-01-02")
	task := Task{
		Date: today,
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
