package converter

import (
	"fmt"
	"reflect"
	"time"
)

type ConvertFunc func(any) (string, bool)

type Converter struct {
	String   func(string) string
	Int      func(int64) string
	UInt     func(uint64) string
	Float    func(float64) string
	Bool     func(bool) string
	Time     func(time.Time) string
	Duration func(time.Duration) string
	Nil      func() string

	ConvertFunc ConvertFunc
	Fallback    func(any) string
}

func (c *Converter) Copy() *Converter {
	return &Converter{
		String:   c.String,
		Int:      c.Int,
		UInt:     c.UInt,
		Float:    c.Float,
		Bool:     c.Bool,
		Time:     c.Time,
		Duration: c.Duration,
		Nil:      c.Nil,
		Fallback: c.Fallback,
	}
}

var DefaultFallback = func(v any) string {
	return fmt.Sprintf("%+v", v)
}

func Default() *Converter {
	return &Converter{
		Float:    FloatDecimalFormatter(2),
		Nil:      NilFormatter,
		Fallback: DefaultFallback,
	}
}

func New() *Converter {
	return Default()
}

func (c *Converter) Convert(v any) string {
	if c.ConvertFunc != nil {
		if s, ok := c.ConvertFunc(v); ok {
			return s
		}
	}

	val := reflect.ValueOf(v)
	if !val.IsValid() {
		if c.Nil != nil {
			return c.Nil()
		} else {
			return c.Fallback(v)
		}
	}

	switch val.Kind() {
	case reflect.String:
		if c.String != nil {
			return c.String(val.String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch v := v.(type) {
		case time.Duration:
			if c.Duration != nil {
				return c.Duration(v)
			}
		default:
			if c.Int != nil {
				return c.Int(val.Int())
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if c.UInt != nil {
			return c.UInt(val.Uint())
		}
	case reflect.Float32, reflect.Float64:
		if c.Float != nil {
			return c.Float(val.Float())
		}
	case reflect.Bool:
		if c.Bool != nil {
			return c.Bool(val.Bool())
		}
	case reflect.Struct:
		switch v := v.(type) {
		case time.Time:
			if c.Time != nil {
				return c.Time(v)
			}
		}
	case reflect.Pointer:
		if val.IsNil() {
			if c.Nil != nil {
				return c.Nil()
			} else {
				return c.Fallback(v)
			}
		}
		return c.Convert(val.Elem().Interface())
	}

	return c.Fallback(v)
}

func FloatDecimalFormatter(precision int) func(float64) string {
	return func(f float64) string {
		return fmt.Sprintf("%.*f", precision, f)
	}
}

func NilFormatter() string {
	return ""
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

func TimeFormatter(format string) func(time.Time) string {
	return func(t time.Time) string {
		return t.Format(format)
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
