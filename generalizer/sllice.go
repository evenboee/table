package generalizer

import (
	"errors"
	"reflect"
)

var (
	ErrNotSlice = errors.New("not a slice")
)

func Slice(v any, opts ...ParamsOption) ([]string, [][]string) {
	return New(opts...).Slice(v)
}

func (params *Params) Slice(v any) ([]string, [][]string) {
	return params.slice_(reflect.ValueOf(v))
}

func (params *Params) slice_(v reflect.Value) ([]string, [][]string) {
	if v.Kind() == reflect.Pointer {
		return params.slice_(v.Elem())
	}

	if v.Kind() != reflect.Slice {
		panic(ErrNotSlice)
	}

	headers := []string{"N", "Value"}
	rows := make([][]string, 0, v.Len())

	for i := 0; i < v.Len(); i++ {
		rows = append(rows, []string{params.Converter.Convert(i + 1), params.Converter.Convert(v.Index(i).Interface())})
	}

	return headers, rows
}
