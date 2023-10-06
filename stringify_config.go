package table

import "github.com/evenboee/go-tabular/generalizer"

type StringifyConfig struct {
	*stringer
	*generalizer.StringerConfig
	Spreadsheet *string
	Offset      int
	Limit       int // if <= 0, no limit
}

var DefaultSpreadsheetHeader = "#"
var DefaultGeneralizerStringerConfig = generalizer.DefaultStringerConfig

func NewStringifyConfig(stringer *stringer) *StringifyConfig {
	return &StringifyConfig{
		stringer:       stringer,
		StringerConfig: DefaultGeneralizerStringerConfig,
		Spreadsheet:    nil,
	}
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
		WithGeneralizer(s.StringerConfig),
		WithSpreadsheetHeader(s.Spreadsheet),
		WithOffset(s.Offset),
		WithLimit(s.Limit),
	}
}

func WithStringifyConfig(s *StringifyConfig) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.stringer = s.stringer
		c.StringerConfig = s.StringerConfig
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

func WithGeneralizer(s *generalizer.StringerConfig) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.StringerConfig = s
	}
}

func WithTimeFormat(format string) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.StringerConfig.Time = generalizer.TimeStringer(format)
	}
}

func WithDecimalPrecision(precision uint16) StringifyConfigOpt {
	return func(c *StringifyConfig) {
		c.StringerConfig.Float = generalizer.FloatDecimalFormatter(int(precision))
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
