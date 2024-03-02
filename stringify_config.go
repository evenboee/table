package table

import "github.com/evenboee/table/generalizer"

type Params struct {
	Style             *TableStyle
	GeneralizerParams *generalizer.Params

	OmitHeader  bool
	OmitPadding bool

	RightAlignNumeric bool

	// Spreadsheet is the row number column header. is nil, no row number column is added
	Spreadsheet *string

	HeaderAlignment alignment
	HeaderFormatter func(string) string
}

var (
	OmitHeader        bool      = false
	GeneralizerParams           = generalizer.Default()
	OmitPadding       bool      = false
	RightAlignNumeric bool      = true
	Spreadsheet       *string   = nil
	HeaderAlignment   alignment = AlignCenter
)

func copy[T any](v *T) *T {
	var vv *T
	if v != nil {
		vv = new(T)
		*vv = *v
	}
	return vv
}

func Default() *Params {
	return &Params{
		Style:             DefaultStyle,
		GeneralizerParams: GeneralizerParams,
		OmitHeader:        OmitHeader,
		OmitPadding:       OmitPadding,
		RightAlignNumeric: RightAlignNumeric,
		Spreadsheet:       copy(Spreadsheet),
		HeaderAlignment:   HeaderAlignment,
	}
}

func New(opts ...ParamsOption) *Params {
	s := Default()

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Params) Copy() *Params {
	return &Params{
		Style:             s.Style.Copy(),
		GeneralizerParams: s.GeneralizerParams.Copy(),
		OmitHeader:        s.OmitHeader,
		OmitPadding:       s.OmitPadding,
		RightAlignNumeric: s.RightAlignNumeric,
		Spreadsheet:       copy(s.Spreadsheet),
	}
}

func (s *Params) With(opts ...ParamsOption) *Params {
	ss := s.Copy()

	for _, opt := range opts {
		opt(ss)
	}

	return ss
}

type ParamsOption func(*Params)

func WithStyle(style *TableStyle) ParamsOption {
	return func(s *Params) {
		s.Style = style
	}
}

func WithGeneralizer(params *generalizer.Params) ParamsOption {
	return func(s *Params) {
		s.GeneralizerParams = params
	}
}

func WithOmitHeader(omit bool) ParamsOption {
	return func(s *Params) {
		s.OmitHeader = omit
	}
}

func WithOmitPadding(omit bool) ParamsOption {
	return func(s *Params) {
		s.OmitPadding = omit
	}
}

func WithRightAlignNumeric(right bool) ParamsOption {
	return func(s *Params) {
		s.RightAlignNumeric = right
	}
}

func WithSpreadsheet(header string) ParamsOption {
	return func(s *Params) {
		s.Spreadsheet = &header
	}
}

func WithHeaderAlignment(align alignment) ParamsOption {
	return func(s *Params) {
		s.HeaderAlignment = align
	}
}

func WithHeaderFormatter(f func(string) string) ParamsOption {
	return func(s *Params) {
		s.HeaderFormatter = f
	}
}

func WithAutoFormatHeader() ParamsOption {
	return func(s *Params) {
		s.HeaderFormatter = AutoFormatHeader
	}
}
