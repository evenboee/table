package generalizer

import (
	"errors"
	"reflect"
	"sort"
	"strings"
)

var ErrDataIsNotMap = errors.New("data is not a map")

func (c *Converter) Map(data any) Result {
	keys := make(map[string]struct{})
	rows := make([]map[string]string, 0)

	v := dereference(reflect.ValueOf(data))
	if v.Kind() != reflect.Slice && v.Elem().Kind() != reflect.Map {
		panic(ErrDataIsNotMap)
	}

	for i := 0; i < v.Len(); i++ {
		m := v.Index(i)
		row := make(map[string]string)

		for _, k := range m.MapKeys() {
			key := c.ToString(k.Interface())
			keyVal := m.MapIndex(k)
			val := c.ToString(keyVal.Interface())

			keys[key] = struct{}{}
			row[key] = val
		}
		rows = append(rows, row)
	}

	headers := make([]string, 0, len(keys))
	for k := range keys {
		headers = append(headers, k)
	}

	sort.Strings(headers)

	return Result{
		Headers: headers,
		Rows:    rows,
	}
}

func (c *Converter) SingleMap(data any) Result {
	rows := make([]map[string]string, 0)

	v := dereference(reflect.ValueOf(data))
	if v.Kind() != reflect.Map {
		panic(ErrDataIsNotMap)
	}

	for _, k := range v.MapKeys() {
		key := c.ToString(k.Interface())
		keyVal := v.MapIndex(k)
		val := c.ToString(keyVal.Interface())

		rows = append(rows, map[string]string{
			"Key":   key,
			"Value": val,
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		return strings.Compare(rows[i]["Key"], rows[j]["Key"]) < 0
	})

	return Result{
		Headers: []string{"Key", "Value"},
		Rows:    rows,
	}
}
