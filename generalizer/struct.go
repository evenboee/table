package generalizer

import (
	"errors"
	"reflect"
	"slices"
)

var (
	ErrNotStruct = errors.New("not a struct")
)

func Struct(v any, opts ...ParamsOption) ([]string, [][]string) {
	return New(opts...).Struct(v)
}

func (params *Params) Struct(v any) ([]string, [][]string) {
	return params.struct_(reflect.ValueOf(v))
}

func (params *Params) struct_(value reflect.Value) ([]string, [][]string) {
	if value.Kind() == reflect.Pointer {
		return params.struct_(value.Elem())
	}

	if value.Kind() != reflect.Slice {
		panic(ErrNotSlice)
	}

	headers := make([]string, 0)
	rows := make([][]string, 0)

	l := value.Len()
	for i := 0; i < l; i++ {
		h, r := params.generalizeStruct(value.Index(i))
		if i == 0 {
			headers = h
		}
		rows = append(rows, r)
	}

	return headers, rows
}

func (params *Params) SingleStruct(v any) ([]string, [][]string) {
	return params.singleStruct_(reflect.ValueOf(v))
}

func (params *Params) singleStruct_(value reflect.Value) ([]string, [][]string) {
	headers := []string{"Field", "Value"}

	fields, values := params.generalizeStruct(value)
	rows := make([][]string, 0, len(values))

	for i, field := range fields {
		rows = append(rows, []string{field, values[i]})
	}

	return headers, rows
}

func (params *Params) generalizeStruct(value reflect.Value) ([]string, []string) {
	if value.Kind() == reflect.Pointer {
		return params.generalizeStruct(value.Elem())
	}

	if value.Kind() != reflect.Struct {
		panic(ErrNotStruct)
	}

	t := value.Type()
	numFields := t.NumField()

	headers := make([]string, 0, numFields)
	row := make([]string, 0, numFields)

	for i := 0; i < numFields; i++ {
		fieldT := t.Field(i)

		if !fieldT.IsExported() {
			continue
		}

		if params.Include != nil {
			if !slices.Contains(params.Include, fieldT.Name) {
				continue
			}
		}

		if params.Exclude != nil {
			if slices.Contains(params.Exclude, fieldT.Name) {
				continue
			}
		}

		tag := fieldT.Tag.Get(params.Tag)
		if tag == "" {
			tag = fieldT.Name
		} else if tag == "-" {
			continue
		}

		headers = append(headers, tag)
		row = append(row, params.Converter.Convert(value.Field(i).Interface()))
	}

	return headers, row
}
