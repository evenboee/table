package table

import "github.com/evenboee/go-tabular/generalizer"

func Array[T any](data []T, opts ...StringifyConfigOpt) string {
	str := NewStringifyConfig(DefaultStringer)
	for _, opt := range opts {
		opt(str)
	}

	return str.Stringify(generalizer.Array(data))
}

func Struct[T any](data []T, opts ...StringifyConfigOpt) string {
	str := NewStringifyConfig(DefaultStringer)
	for _, opt := range opts {
		opt(str)
	}

	return str.Stringify(generalizer.Struct(data, str.StringerConfig))
}

func StructLimit[T any](data []T, limit int, opts ...StringifyConfigOpt) []string {
	str := NewStringifyConfig(DefaultStringer)
	for _, opt := range opts {
		opt(str)
	}

	str.Limit = limit
	if str.Limit <= 0 {
		str.Limit = len(data)
	}

	headers, rows := generalizer.Struct(data, str.StringerConfig)

	groups := make([]string, 0, len(data)/str.Limit)

	for i := str.Offset; i < len(data); i += str.Limit {
		str.Offset = i
		s := str.Stringify(headers, rows)
		groups = append(groups, s)
	}

	return groups
}

func SingleStruct[T any](data T, opts ...StringifyConfigOpt) string {
	str := NewStringifyConfig(DefaultStringer)
	for _, opt := range opts {
		opt(str)
	}

	return str.Stringify(generalizer.SingleStruct(data, str.StringerConfig))
}

func Map[K comparable, V any](data []map[K]V, order []K, opts ...StringifyConfigOpt) string {
	str := NewStringifyConfig(DefaultStringer)
	for _, opt := range opts {
		opt(str)
	}

	return str.Stringify(generalizer.Map(data, order))
}
