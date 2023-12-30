package generalizer

import (
	"reflect"
	"strconv"
)

func (c *Converter) Array(data any) Result {
	headers := []string{"N", "Value"}

	rows := make([]map[string]string, 0)
	v := dereference(reflect.ValueOf(data))
	if v.Kind() != reflect.Slice {
		panic("data is not a slice")
	}

	for i := 0; i < v.Len(); i++ {
		rows = append(rows, map[string]string{
			"N":     strconv.Itoa(i),
			"Value": c.ToString(v.Index(i).Interface()),
		})
	}

	return Result{
		Headers: headers,
		Rows:    rows,
	}
}
