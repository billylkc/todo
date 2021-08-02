package todo

// WipeAll removes all the tasks
func WipeAll(path string) error {

	var tasks []Task
	err := WriteTasks(path, tasks)
	if err != nil {
		return err
	}
	return nil
}
