package generalizer

import (
	"github.com/evenboee/table/converter"
)

type Params struct {
	Converter *converter.Converter

	// (struct) tag key to get header value from
	Tag string
	// (struct) if not empty, include only specified fields
	Include []string
	// (struct) if not empty, exclude no fields. applied after Include
	Exclude []string
}

var DefaultTag = "table"

func Default() *Params {
	return &Params{
		Converter: converter.Default(),
		Tag:       DefaultTag,
	}
}

func New(opts ...ParamsOption) *Params {
	params := Default()
	for _, opt := range opts {
		opt(params)
	}
	return params
}

func (params *Params) Copy() *Params {
	return &Params{
		Converter: params.Converter.Copy(),
		Tag:       params.Tag,
		Include:   append([]string(nil), params.Include...),
		Exclude:   append([]string(nil), params.Exclude...),
	}
}

type ParamsOption func(*Params)

func Converter(opts ...converter.ConverterOption) ParamsOption {
	return func(params *Params) {
		params.Converter = params.Converter.With(opts...)
	}
}

func WithConverter(c *converter.Converter) ParamsOption {
	return func(params *Params) {
		params.Converter = c
	}
}

func WithTag(tag string) ParamsOption {
	return func(params *Params) {
		params.Tag = tag
	}
}

func Include(include ...string) ParamsOption {
	return func(params *Params) {
		params.Include = include
	}
}

func Exclude(exclude ...string) ParamsOption {
	return func(params *Params) {
		if len(exclude) == 0 {
			exclude = nil
		}
		params.Exclude = exclude
	}
}
