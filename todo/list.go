package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ListTasks lists all the tasks
func ListTasks(path string, tag string, undoneOnly bool) error {

	tag = strings.ToLower(strings.TrimSpace(tag))

	// Read from file
	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Print result
	fmt.Printf("\nid\tdate\t\tdone\ttodo\n")
	fmt.Printf("--\t----\t\t----\t----\n")
	for i, task := range tasks {
		line := fmt.Sprintf("%d\t%s\t[%s]\t%s", i+1, task.Date, task.Done, task.Name)

		// Status filtering
		if undoneOnly {
			if task.Done == "x" {
				continue // pass for done tasks
			}
		}

		// Tag filtering
		if tag != "" {
			if strings.Contains(strings.ToLower(line), tag) {
				// print
			} else {
				continue //skip
			}
		}
		fmt.Println(line)
	}
	fmt.Println("")

	return nil
}

// ReadTasks reads all the tasks from the todofile
func ReadTasks(path string) ([]Task, error) {
	var tasks []Task

	file, err := os.Open(path)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	if err != nil {
		return tasks, err
	}

	// Read file and construct struct
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ",")
		if len(splits) == 3 {
			date := strings.TrimSpace(splits[0])
			name := strings.TrimSpace(splits[1])
			done := strings.TrimSpace(splits[2])

			// handle done
			if done != "x" {
				done = " "
			}

			task := Task{
				Date: date,
				Name: name,
				Done: done,
			}
			tasks = append(tasks, task)
		} else {
			// continue
		}
	}
	if err := scanner.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}
