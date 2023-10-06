package generalizer

import (
	"fmt"
	"time"
)

type StringerConfig struct {
	String func(string) string
	Bool   func(bool) string
	Int    func(int64) string
	Float  func(float64) string
	Time   func(time.Time) string
	Nil    func() string
}

var DefaultStringerConfig = &StringerConfig{
	Time:  TimeStringer(time.RFC3339),
	Float: FloatDecimalFormatter(2),
	Nil:   NilStringer,
}

func NewStringerConfig() *StringerConfig {
	return &StringerConfig{}
}

func (s *StringerConfig) ToString(v any) string {
	if s == nil || v == nil {
		if s.Nil != nil {
			return s.Nil()
		}
		return ""
	}

	switch v := v.(type) {
	case string:
		if s.String != nil {
			return s.String(v)
		}
	case bool:
		if s.Bool != nil {
			return s.Bool(v)
		}
	case int, int8, int16, int32, int64:
		if s.Int != nil {
			if v, ok := v.(int64); ok {
				return s.Int(v)
			}
		}
	case float64, float32:
		if s.Float != nil {
			if v, ok := v.(float64); ok {
				return s.Float(v)
			}
		}
	case time.Time:
		if s.Time != nil {
			return s.Time(v)
		}
	}

	return fmt.Sprintf("%v", v)
}

func TimeStringer(format string) func(time.Time) string {
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

func NilStringer() string {
	return ""
}
