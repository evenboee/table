package generalizer

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestConverter_ToString(t *testing.T) {
	type pair struct {
		val  any
		want string
	}

	tc := []struct {
		name  string
		c     func() *Converter
		pairs []pair
	}{
		{
			name: "nil",
			c: func() *Converter {
				conv := DefaultConverter
				conv.Nil = func() string {
					return "null"
				}
				return conv
			},
			pairs: []pair{
				{nil, "null"},
				{(*int)(nil), "null"},
			},
		},
		{
			name: "string",
			c: func() *Converter {
				conv := DefaultConverter
				conv.String = func(s string) string {
					return s + s
				}
				return conv
			},
			pairs: []pair{
				{"hello", "hellohello"},
				{"", ""},
			},
		},
		{
			name: "bool",
			c: func() *Converter {
				conv := DefaultConverter
				conv.Bool = func(b bool) string {
					if b {
						return "yes"
					}
					return "no"
				}
				return conv
			},
			pairs: []pair{
				{true, "yes"},
				{false, "no"},
			},
		},
		{
			name: "int",
			c: func() *Converter {
				conv := DefaultConverter
				conv.Int = func(i int64) string {
					return strconv.Itoa(int(i * 2))
				}
				conv.UInt = func(i uint64) string {
					return strconv.Itoa(int(i * 2))
				}
				return conv
			},
			pairs: []pair{
				{1, "2"},
				{0, "0"},
				{-1, "-2"},
				{int8(2), "4"},
				{int16(3), "6"},
				{int32(4), "8"},
				{int64(5), "10"},
				{uint(6), "12"},
				{uint8(7), "14"},
			},
		},
		{
			name: "float",
			c: func() *Converter {
				conv := DefaultConverter
				conv.Float = func(f float64) string {
					return fmt.Sprintf("%.2f", f*2)
				}
				return conv
			},
			pairs: []pair{
				{1.123456, "2.25"},
				{0.0, "0.00"},
				{-1.123456, "-2.25"},
				{float32(1.123456), "2.25"},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c()
			for _, p := range tt.pairs {
				got := c.ToString(p.val)
				require.Equal(t, p.want, got)
			}
		})
	}
}
