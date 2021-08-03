package todo

import (
	"bufio"
	"fmt"
	"os"
)

// BackupTasks reads from the original task file and append the content to the backup file
func BackupTasks(path string) error {
	var lines []string

	// Read the files in todo
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if len(lines) == 0 {
		return fmt.Errorf("/nNo new todo items for archive. /n")
	}

	// Append to backup files
	backupPath := path + ".bk" // ~/todo.txt.bk
	bk, err := os.OpenFile(backupPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(bk)
	if len(lines) > 0 {
		for _, line := range lines {
			_, _ = datawriter.WriteString(line + "\n")
		}
	}
	datawriter.Flush()
	if err := bk.Close(); err != nil {
		return err
	}
	fmt.Printf("\nWritten %d lines to backup file [%s]\n\n", len(lines), backupPath)

	return nil
}
