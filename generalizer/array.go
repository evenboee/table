package generalizer

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrDataIsNotSlice = errors.New("data is not a slice")

func (c *Converter) Array(data any) Result {
	headers := []string{"N", "Value"}

	rows := make([]map[string]string, 0)
	v := dereference(reflect.ValueOf(data))
	if v.Kind() != reflect.Slice {
		panic(ErrDataIsNotSlice)
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
