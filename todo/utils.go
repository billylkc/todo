package todo

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertID converts the input id to int
func ConvertID(args []string) (int, error) {
	var (
		id  int
		err error
	)
	if len(args) == 1 {
		id, err = strconv.Atoi(args[0])
		if err != nil {
			return id, fmt.Errorf("Invalid task ID input - %s \n", args[1])
		}
	} else {
		return id, fmt.Errorf("Invalid task ID input - %v \n\n", strings.Join(args, ", "))
	}
	return id, nil
}
