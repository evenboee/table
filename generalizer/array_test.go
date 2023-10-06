package generalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumericArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expectedHeaders := []string{"N", "Value"}
	expectedRows := []map[string]string{
		{
			"N":     "0",
			"Value": "1",
		}, {
			"N":     "1",
			"Value": "2",
		}, {
			"N":     "2",
			"Value": "3",
		}, {
			"N":     "3",
			"Value": "4",
		}, {
			"N":     "4",
			"Value": "5",
		},
	}

	headers, rows := Array(input)
	require.Equal(t, expectedHeaders, headers)
	require.Equal(t, expectedRows, rows)
}

func TestNonNumericArray(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e"}
	expectedHeaders := []string{"N", "Value"}
	expectedRows := []map[string]string{
		{
			"N":     "0",
			"Value": "a",
		}, {
			"N":     "1",
			"Value": "b",
		}, {
			"N":     "2",
			"Value": "c",
		}, {
			"N":     "3",
			"Value": "d",
		}, {
			"N":     "4",
			"Value": "e",
		},
	}

	headers, rows := Array(input)
	require.Equal(t, expectedHeaders, headers)
	require.Equal(t, expectedRows, rows)
}

func TestAnyArray(t *testing.T) {
	input := []any{1, "a", 2, "b", 3, "c"}
	expectedHeaders := []string{"N", "Value"}
	expectedRows := []map[string]string{
		{
			"N":     "0",
			"Value": "1",
		}, {
			"N":     "1",
			"Value": "a",
		}, {
			"N":     "2",
			"Value": "2",
		}, {
			"N":     "3",
			"Value": "b",
		}, {
			"N":     "4",
			"Value": "3",
		}, {
			"N":     "5",
			"Value": "c",
		},
	}

	headers, rows := Array(input)
	require.Equal(t, expectedHeaders, headers)
	require.Equal(t, expectedRows, rows)
}
