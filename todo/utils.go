package todo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/now"
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

// ParseInputDate parse the task date into right format
// Support now package and mmdd, dd/mmformat
func ParseInputDate(in string) (string, error) {

	// Handle mmdd format
	if len(in) == 4 {
		in = in[0:2] + "-" + in[2:]
	}
	// Handle dd/mm format
	if len(in) == 5 && strings.Contains(in, "/") {
		in = in[3:] + "-" + in[0:2]
	}
	t, err := now.Parse(in)
	if err != nil {
		return "", err
	}

	out := t.Format("2006-01-02")
	return out, nil
}
