package generalizer

import (
	"reflect"
	"testing"
)

func requireEqual(t *testing.T, expected any, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func requireTable(t *testing.T,
	expectedHeaders []string, expectedRows [][]string,
	actualHeaders []string, actualRows [][]string,
) {
	t.Helper()
	requireEqual(t, expectedHeaders, actualHeaders)
	requireEqual(t, expectedRows, actualRows)
}

func testTable(t *testing.T,
	data any,
	expectedHeaders []string, expectedRows [][]string,
	opts ...ParamsOption,
) {
	t.Helper()
	params := New(opts...)
	headers, rows := params.Any(data)
	requireTable(t, expectedHeaders, expectedRows, headers, rows)
}

func requirePanic(t *testing.T, f func(), expected any) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic")
		} else {
			requireEqual(t, expected, r)
		}
	}()
	f()
}
