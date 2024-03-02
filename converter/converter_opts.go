package converter

import (
	"fmt"
	"time"
)

type ConverterOption func(*Converter)

func (c *Converter) Set(f func(c *Converter)) *Converter {
	f(c)
	return c
}

func (c *Converter) With(opts ...ConverterOption) *Converter {
	cc := c.Copy()

	for _, opt := range opts {
		opt(cc)
	}
	return cc
}

func String(f func(string) string) ConverterOption {
	return func(c *Converter) {
		c.String = f
	}
}

func Int(f func(int64) string) ConverterOption {
	return func(c *Converter) {
		c.Int = f
	}
}

func UInt(f func(uint64) string) ConverterOption {
	return func(c *Converter) {
		c.UInt = f
	}
}

func Float(precision int) ConverterOption {
	if precision < 0 {
		return func(c *Converter) {
			c.Float = func(f float64) string {
				return fmt.Sprintf("%f", f)
			}
		}
	}

	return func(c *Converter) {
		c.Float = func(f float64) string {
			return fmt.Sprintf("%.*f", precision, f)
		}
	}
}

func Bool(t string, f string) ConverterOption {
	return func(c *Converter) {
		c.Bool = BoolFormatter(t, f)
	}
}

func Time(format string) ConverterOption {
	return func(c *Converter) {
		c.Time = TimeFormatter(format)
	}
}

func Duration(f func(time.Duration) string) ConverterOption {
	return func(c *Converter) {
		c.Duration = f
	}
}

func Nil(n string) ConverterOption {
	return func(c *Converter) {
		c.Nil = func() string {
			return n
		}
	}
}
