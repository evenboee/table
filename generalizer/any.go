package generalizer

import (
	"errors"
	"reflect"
)

var ErrUnsupportedType = errors.New("unsupported type")

// Any is a router for other functions.
// It will determine the type of the data and call the appropriate generalizer function.
// If the type is not supported, it will panic with ErrUnsupportedType.
//
// Supported types are: []struct, []map, []any, struct, map
func (c *Converter) Any(data any) Result {
	if data == nil {
		return Result{
			Headers: make([]string, 0),
			Rows:    make([]map[string]string, 0),
		}
	}

	t := dereference(reflect.TypeOf(data))

	if t.Kind() == reflect.Slice {
		switch t.Elem().Kind() {
		case reflect.Struct:
			return c.Struct(data)
		case reflect.Map:
			return c.Map(data)
		default:
			return c.Array(data)
		}
	} else if t.Kind() == reflect.Struct {
		return c.SingleStruct(data)
	} else if t.Kind() == reflect.Map {
		return c.SingleMap(data)
	} else {
		panic(ErrUnsupportedType)
	}

	return Result{}
}
