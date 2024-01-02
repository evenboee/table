package table

import "github.com/evenboee/table/generalizer"

type StringifyConfig struct {
	*stringer
	*generalizer.Converter
	Spreadsheet *string
	Offset      int
	Limit       int // if <= 0, no limit
}

var DefaultSpreadsheetHeader = "#"

func NewStringifyConfig(stringer *stringer) *StringifyConfig {
	return &StringifyConfig{
		stringer:    stringer,
		Converter:   generalizer.GetDefaultConverter(),
		Spreadsheet: nil,
	}
}

func DefaultStringifyConfig() *StringifyConfig {
	return NewStringifyConfig(DefaultStringer)
}

type StringifyConfigOpt func(*StringifyConfig)

func (s *StringifyConfig) With(opts ...StringifyConfigOpt) *StringifyConfig {
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *StringifyConfig) GetOpts() []StringifyConfigOpt {
	return []StringifyConfigOpt{
		WithStringer(s.stringer),
		WithConverter(s.Converter),
		WithSpreadsheetHeader(s.Spreadsheet),
		WithOffset(s.Offset),
		WithLimit(s.Limit),
	}
}

func WithStringifyConfig(s *StringifyConfig) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.stringer = s.stringer
		c.Converter = s.Converter
		c.Spreadsheet = s.Spreadsheet
		c.Offset = s.Offset
		c.Limit = s.Limit
	}
}

func WithStringer(s *stringer) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.stringer = s
	}
}

func WithDefaultSpreadsheet(c *StringifyConfig) {
	c.Spreadsheet = &DefaultSpreadsheetHeader
}

func WithSpreadsheet(s string) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Spreadsheet = &s
	}
}

func WithConverter(s *generalizer.Converter) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Converter = s
	}
}

func WithTimeFormat(format string) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Converter.Time = generalizer.TimeFormatter(format)
	}
}

func WithDecimalPrecision(precision uint16) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Converter.Float = generalizer.FloatDecimalFormatter(int(precision))
	}
}

func WithNilFormatter(s string) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Converter.Nil = generalizer.NilAsFormatter(s)
	}
}

func WithBoolFormatter(t string, f string) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Converter.Bool = generalizer.BoolFormatter(t, f)
	}
}

func WithSpreadsheetHeader(s *string) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Spreadsheet = s
	}
}

func WithOffset(offset int) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Offset = offset
	}
}

func WithLimit(limit int) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.Limit = limit
	}
}
