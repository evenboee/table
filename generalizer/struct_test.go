package generalizer

import (
	"reflect"
	"testing"
)

func Test__struct_(t *testing.T) {
	testTable(t, []struct {
		A           int
		B           string
		c           float64
		D           string `table:"-"`
		E           string `table:"Test"`
		Excluded    string
		NotIncluded string
	}{
		{1, "one", 3.14, "four", "five", "six", "seven"},
		{2, "two", 6.28, "eight", "nine", "ten", "eleven"},
	}, []string{"A", "B", "Test"}, [][]string{
		{"1", "one", "five"},
		{"2", "two", "nine"},
	}, Exclude("Excluded"), Include("A", "B", "c", "D", "E", "Excluded"))

	// panic (not slice)
	requirePanic(t, func() {
		Default().struct_(reflect.ValueOf(1))
	}, ErrNotSlice)
}

func Test__generalizeStruct(t *testing.T) {
	headers, rows := Default().generalizeStruct(reflect.ValueOf(&struct {
		A string
		B int `table:"BB"`
		C float64
	}{"hello", 123, 3.14}))
	requireEqual(t, []string{"A", "BB", "C"}, headers)
	requireEqual(t, []string{"hello", "123", "3.14"}, rows)

	// panic
	requirePanic(t, func() {
		Default().generalizeStruct(reflect.ValueOf(1))
	}, ErrNotStruct)
}

func Test__singleStruct_(t *testing.T) {
	testTable(t, struct {
		A string
		B int `table:"BB"`
		C float64
	}{"hello", 123, 3.14}, []string{"Field", "Value"}, [][]string{
		{"A", "hello"},
		{"BB", "123"},
		{"C", "3.14"},
	})
}
