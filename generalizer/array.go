package generalizer

import (
	"fmt"
	"strconv"
)

func Array[T any](data []T) ([]string, []map[string]string) {
	headers := []string{"N", "Value"}

	rows := make([]map[string]string, len(data))
	for i, row := range data {
		rows[i] = map[string]string{
			"N":     strconv.Itoa(i),
			"Value": fmt.Sprintf("%v", row),
		}
	}

	return headers, rows
}
