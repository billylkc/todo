package todo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/now"
)

// ConvertID converts the input id to int
// Handle multiple values with format 2-4, 2,3,4
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

// ConvertID converts the input id to []int. Similar to ConvertID
// Handle multiple values with format 2-4, 2,3,4
func ConvertIDs(args []string) ([]int, error) {
	var (
		ids []int
		err error
	)
	if len(args) == 1 {

		in := args[0]

		// Handle format like 2-4
		if strings.Contains(in, "-") {
			splits := strings.Split(in, "-")
			if len(splits) == 2 {

				var one, two int
				num1 := splits[0]
				num2 := splits[1]

				one, err = strconv.Atoi(num1)
				if err != nil {
					return ids, err
				}
				two, err = strconv.Atoi(num2)
				if err != nil {
					return ids, err
				}
				numRange := makeRange(one, two)
				return numRange, nil

			} else {
				return ids, fmt.Errorf("Invalid task ID input - %s \n", args[0])
			}
		}

		// Handle format like 2,3,4
		if strings.Contains(in, ",") {
			splits := strings.Split(in, ",")
			for _, c := range splits {
				var id int
				id, err = strconv.Atoi(c)
				if err != nil {
					return ids, fmt.Errorf("Invalid task ID input - %s \n", args[0])
				}
				ids = append(ids, id)
			}
			return ids, nil
		}

		// Normal case, single id
		id, err := strconv.Atoi(in)
		if err != nil {
			return ids, fmt.Errorf("Invalid task ID input - %s \n", args[0])
		}
		ids = []int{id}

	} else {
		return ids, fmt.Errorf("Invalid task ID input - %v \n\n", strings.Join(args, ", "))
	}
	return ids, nil
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

// makeRange generates a list of integer of min, max input (inclusive)
func makeRange(min, max int) []int {
	// swap if min, max are not valid
	if max < min {
		max, min = min, max
	}

	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
