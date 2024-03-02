package generalizer

import (
	"reflect"
	"testing"
)

func Test__any_(t *testing.T) {
	params := Default()
	nilValue := "<<nil>>"
	params.Converter.Nil = func() string { return nilValue }

	// invalid
	headers, rows := params.any_(reflect.ValueOf(nil))
	requireTable(t, []string{"Key", "Value"}, [][]string{
		{"Value", nilValue},
		{"Type", "nil"},
	}, headers, rows)

	// nil simple value + pointer recursive
	var nilInt *int
	headers, rows = params.any_(reflect.ValueOf(&nilInt))
	requireTable(t, []string{"Key", "Value"}, [][]string{
		{"Value", nilValue},
		{"Type", "*int"},
	}, headers, rows)
}
