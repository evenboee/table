package generalizer

import (
	"errors"
	"reflect"
	"slices"
)

var (
	ErrNotMap = errors.New("not a map")
)

func Map(v any, opts ...ParamsOption) ([]string, [][]string) {
	return New(opts...).Map(v)
}

func (params *Params) Map(v any) ([]string, [][]string) {
	return params.map_(reflect.ValueOf(v))
}

func (params *Params) map_(value reflect.Value) ([]string, [][]string) {
	if value.Kind() == reflect.Pointer {
		return params.struct_(value.Elem())
	}

	if value.Kind() != reflect.Slice {
		panic(ErrNotSlice)
	}

	headersSet := make(map[string]struct{}, 0)
	rowsSet := make([]map[string]string, 0)

	l := value.Len()
	for i := 0; i < l; i++ {
		h, r := params.generalizeMap(value.Index(i))
		for _, header := range h {
			headersSet[header] = struct{}{}
		}

		rowsSet = append(rowsSet, r)
	}

	headers := make([]string, 0, len(headersSet))
	for header := range headersSet {
		headers = append(headers, header)
	}

	slices.Sort(headers)

	rows := make([][]string, 0, len(rowsSet))
	for _, row := range rowsSet {
		rowSlice := make([]string, 0, len(headers))
		for _, header := range headers {
			rowSlice = append(rowSlice, row[header])
		}
		rows = append(rows, rowSlice)
	}

	return headers, rows
}

func (params *Params) SingleMap(v any) ([]string, [][]string) {
	return params.singleMap_(reflect.ValueOf(v))
}

func (params *Params) singleMap_(value reflect.Value) ([]string, [][]string) {
	keys, values := params.generalizeMap(value)
	headers := []string{"Key", "Value"}
	rows := make([][]string, 0, len(keys))

	slices.Sort(keys)

	for _, key := range keys {
		rows = append(rows, []string{key, values[key]})
	}

	return headers, rows
}

func (params *Params) generalizeMap(value reflect.Value) ([]string, map[string]string) {
	if value.Kind() == reflect.Pointer {
		return params.generalizeMap(value.Elem())
	}

	if value.Kind() != reflect.Map {
		panic(ErrNotMap)
	}

	mapKeys := value.MapKeys()
	keys := make([]string, 0, len(mapKeys))
	values := make(map[string]string, len(mapKeys))

	for _, key := range mapKeys {
		k := params.Converter.Convert(key.Interface())
		v := params.Converter.Convert(value.MapIndex(key).Interface())

		keys = append(keys, k)
		values[k] = v
	}

	return keys, values
}
