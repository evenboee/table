package generalizer

import (
	"reflect"
	"testing"
)

func Test__map_(t *testing.T) {
	testTable(t, []map[int]string{
		{1: "one", 2: "two", 3: "three"},
		{1: "one2", 2: "two2", 3: "three2", 4: "four2"},
	}, []string{"1", "2", "3", "4"}, [][]string{
		{"one", "two", "three", ""},
		{"one2", "two2", "three2", "four2"},
	})

	// panic (not slice)
	requirePanic(t, func() {
		Default().map_(reflect.ValueOf(1))
	}, ErrNotSlice)

}

func Test__singleMap_(t *testing.T) {
	testTable(t, map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}, []string{"Key", "Value"}, [][]string{
		{"1", "one"},
		{"2", "two"},
		{"3", "three"},
	})

}

func Test__generalizeMap(t *testing.T) {
	keys, values := Default().generalizeMap(reflect.ValueOf(&map[string]string{"a": "b"}))
	requireEqual(t, []string{"a"}, keys)
	requireEqual(t, map[string]string{"a": "b"}, values)

	// panic
	requirePanic(t, func() {
		Default().generalizeMap(reflect.ValueOf(1))
	}, ErrNotMap)
}
