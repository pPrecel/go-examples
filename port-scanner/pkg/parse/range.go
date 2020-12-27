package parse

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	portSeparator  = ","
	rangeSeparator = "-"
)

func Range(input string) ([]int, error) {
	sepVal := strings.Split(input, portSeparator)
	var newRange []int

	for _, val := range sepVal {
		if strings.Contains(val, rangeSeparator) {
			minMax := strings.Split(val, rangeSeparator)
			if len(minMax) != 2 {
				return newRange, fmt.Errorf("incorrect format of %s", val)
			}

			min, err := strconv.Atoi(minMax[0])
			if err != nil {
				return newRange, err
			}
			max, err := strconv.Atoi(minMax[1])
			if err != nil {
				return newRange, err
			}
			if min > max {
				return newRange, fmt.Errorf("incorrect format of %s", val)
			}

			for i := min; i <= max; i++ {
				newRange = append(newRange, i)
			}
		} else {
			i, err := strconv.Atoi(val)
			if err != nil {
				return newRange, err
			}

			newRange = append(newRange, i)
		}
	}

	return newRange, nil
}
