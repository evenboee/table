package generalizer

import (
	"reflect"
	"testing"
)

func Test__slice_(t *testing.T) {
	testTable(t, []int{1, 2, 3}, []string{"N", "Value"}, [][]string{
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
	})

	headers, rows := Default().slice_(reflect.ValueOf(&[]int{12, 34, 56}))
	requireTable(t, []string{"N", "Value"}, [][]string{
		{"1", "12"},
		{"2", "34"},
		{"3", "56"},
	}, headers, rows)

	// panic (not slice)
	requirePanic(t, func() {
		Default().slice_(reflect.ValueOf(1))
	}, ErrNotSlice)
}
