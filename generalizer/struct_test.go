package generalizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStruct(t *testing.T) {

}

func TestNonStruct(t *testing.T) {
	require.Panics(t, func() {
		_, _ = DefaultConverter.Struct([]int{1, 2, 3})
	})
}
