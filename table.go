package table

// ToString is a wrapper for StringifyConfig.Stringify(StringifyConfig.Converter.Any(data))
func (s *StringifyConfig) ToString(data any) string {
	return s.Stringify(s.Any(data))
}

// String creates a StringifyConfig with the options (opts) provided and the default stringer
// Calls ToString on the created StringifyConfig
func String(data any, opts ...StringifyConfigOpt) string {
	return DefaultStringifyConfig().With(opts...).ToString(data)
}

// StringWith creates a StringifyConfig with the options (opts) provided and the provided stringer
// Calls ToString on the created StringifyConfig
func StringWith(stringer *stringer, data any, opts ...StringifyConfigOpt) string {
	return NewStringifyConfig(stringer).With(opts...).ToString(data)
}
