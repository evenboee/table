package generalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	data := []map[string]int{
		{
			"foo": 1,
			"bar": 2,
		},
		{
			"foo": 3,
			"bar": 4,
		},
		{
			"taz": 5,
		},
	}

	order := []string{"bar", "foo", "zaz"}

	headers, rows := Map(data, order)

	require.Equal(t, 4, len(headers))
	require.Equal(t, []string{"bar", "foo", "zaz", "taz"}, headers)

	require.Equal(t, len(rows), 3)
}
