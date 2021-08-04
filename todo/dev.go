package todo

import (
	"fmt"
	"strconv"
	"strings"
)

// Dev for development purposes
func Dev() error {
	args := []string{"5"}
	res, err := converts(args)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

func converts(args []string) ([]int, error) {
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
