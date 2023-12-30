package generalizer

import (
	"fmt"
	"reflect"
	"time"
)

// The SubConverter interface is used to convert a value to string.
// This is used to stringify other types not otherwise supported by the converter.
type SubConverter interface {
	// ToString returns the string representation of the value.
	// Returns false if the value is not supported by the converter.
	ToString(any) (string, bool)
}

type Converter struct {
	String func(string) string
	Bool   func(bool) string
	Int    func(int64) string
	UInt   func(uint64) string
	Float  func(float64) string
	Time   func(time.Time) string
	Nil    func() string

	Sub SubConverter
}

func GetDefaultConverter() *Converter {
	return &Converter{
		Float: FloatDecimalFormatter(2),
		Time:  TimeFormatter(time.RFC3339),
		Nil:   NilFormatter,
	}
}

func (c *Converter) With(f func(*Converter)) *Converter {
	f(c)
	return c
}

var DefaultConverter = GetDefaultConverter()

// ToString returns the string representation of the value.
// Starts by checking if the SubConverter is set and tries to use it to convert the value.
// Then the value is checked for nil and if the converter has a Nil function, it is used.
// Then there is a type switch for the value and the corresponding function is used if exists.
// If no function is found, fmt.Sprintf is used with %v (uses String() method if exists).
func (c *Converter) ToString(val any) string {
	// Try to use SubConverter
	if c.Sub != nil {
		s, ok := c.Sub.ToString(val)
		if ok {
			return s
		}
	}

	// Check for nil
	if val == nil && c.Nil != nil {
		return c.Nil()
	}

	// Get actual value (no pointer)
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Pointer && v.IsNil() {
		return c.Nil()
	}

	// Check type
	switch val := val.(type) {
	case string:
		if c.String != nil {
			return c.String(val)
		}
	case bool:
		if c.Bool != nil {
			return c.Bool(val)
		}
	case int, int8, int16, int32, int64:
		if c.Int != nil {
			return c.Int(v.Int())
		}
	case uint, uint8, uint16, uint32, uint64:
		if c.UInt != nil {
			return c.UInt(v.Uint())
		}
	case float64, float32:
		if c.Float != nil {
			return c.Float(v.Float())
		}
	case time.Time:
		if c.Time != nil {
			return c.Time(val)
		}
	}

	// Fallback to fmt.Sprintf
	return fmt.Sprintf("%v", val)
}

func TimeFormatter(format string) func(time.Time) string {
	return func(t time.Time) string {
		return t.Format(format)
	}
}

func FloatDecimalFormatter(decimal int) func(float64) string {
	p := fmt.Sprintf("%%.%df", decimal)
	return func(f float64) string {
		return fmt.Sprintf(p, f)
	}
}

func NilFormatter() string {
	return ""
}

func NilAsFormatter(s string) func() string {
	return func() string {
		return s
	}
}

func BoolFormatter(t string, f string) func(bool) string {
	return func(b bool) string {
		if b {
			return t
		}
		return f
	}
}

func StringCutoffFormatter(limit int) func(string) string {
	return func(s string) string {
		if len(s) > limit {
			if limit > 3 {
				return s[:limit-3] + "..."
			} else {
				return s[:limit]
			}
		}
		return s
	}
}

type subConverter struct {
	f func(any) (string, bool)
}

func (c *subConverter) ToString(val any) (string, bool) {
	return c.f(val)
}

func SubConverterFrom(f func(any) (string, bool)) SubConverter {
	return &subConverter{f}
}
