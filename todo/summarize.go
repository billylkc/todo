package todo

import (
	"fmt"
	"sort"
	"strings"
)

// SummarizeTask summarizes the man-day used for different tasks
// Tag is used as a search terms to filter the result
// TODO: bygroup is used to print the result by project group first. Now use [xxx] format.
func SummarizeTask(path string, tag string, bygroup bool) error {

	tag = strings.ToLower(strings.TrimSpace(tag))
	_ = bygroup

	tasks, err := ReadTasks(path)
	if err != nil {
		return err
	}

	// Remove additional information for tasks name (ignore after hypen)
	for i, _ := range tasks {
		tasks[i].Name = trimStringFromChar(tasks[i].Name, "-")
	}

	// Derive mappings for each day / task
	dateMap := make(map[string][]string) // dateMap - key: date, value: [tasks]. Later calulate weightings per day
	taskMap := make(map[string][]string) // taskMap - key: task, value: [dates]. Later sum up weightings per day
	doneMap := make(map[string]string)   // doneMap - key: task, value: [done]. If at least one record is done, then done
	for _, task := range tasks {

		// dateMap
		if val, ok := dateMap[task.Date]; ok {
			dateMap[task.Date] = append(val, task.Name)
		} else {
			dateMap[task.Date] = []string{task.Name}
		}

		// taskMap
		if val, ok := taskMap[task.Name]; ok {
			taskMap[task.Name] = append(val, task.Date)
		} else {
			taskMap[task.Name] = []string{task.Date}
		}

		// doneMap
		if val, ok := doneMap[task.Name]; ok {
			if val == "x" {
				// skip. Already done
			} else {
				doneMap[task.Name] = task.Done
			}
		} else {
			doneMap[task.Name] = task.Done
		}

	}

	// Derive date weightings for each day
	var dates []string // get list of unique dates
	dateWeight := make(map[string]float64)
	for k, v := range dateMap {
		var weight float64
		if len(v) > 0 {
			weight = 1 / float64(len(v))
		}
		dateWeight[k] = weight
		dates = append(dates, k)
	}

	// Derive man-day for each task
	taskWeight := make(map[string]float64)
	for k, v := range taskMap {
		var weight float64
		if len(v) > 0 {
			for _, day := range v {
				weight = weight + dateWeight[day]
			}
		}
		taskWeight[k] = weight
	}

	// Sort task by man-day before display
	keys := make([]string, 0, len(taskWeight))
	for key := range taskWeight {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return taskWeight[keys[i]] > taskWeight[keys[j]] })

	// Display result
	fmt.Printf("\nid\ttime\t\tdone\ttodo\n")
	fmt.Printf("--\t----\t\t----\t----\n")
	for i, key := range keys {
		j := i + 1
		line := fmt.Sprintf("%d\t%.1f days\t [%s]\t%s\n",
			j,
			taskWeight[keys[i]],
			doneMap[keys[i]],
			key)

		if tag != "" {
			if strings.Contains(strings.ToLower(line), tag) {
				// print
			} else {
				continue //skip
			}
		}
		fmt.Printf(line)
	}
	sort.Strings(dates)
	fmt.Printf("\nTotal of %d days\n    %v\n\n", len(dates), dates)
	return nil
}

// trimStringFromChar returns the first part of the string before the provided character
// Used to ignore string after hypen, when those are additonal information, e.g. Meeting - Something
func trimStringFromChar(s string, char string) string {
	var out string
	if idx := strings.Index(s, char); idx != -1 {
		out = s[:idx]
	} else {
		out = s
	}
	out = strings.TrimSpace(out)
	return out
}
