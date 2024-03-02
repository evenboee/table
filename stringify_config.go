package table

import "github.com/evenboee/table/generalizer"

// TODO: Rename to Stringifier
type Stringify struct {
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
	HeaderAlignment   alignment = Center
)

func copy[T any](v *T) *T {
	var vv *T
	if v != nil {
		vv = new(T)
		*vv = *v
	}
	return vv
}

func Default() *Stringify {
	return &Stringify{
		Style:             DefaultStyle,
		GeneralizerParams: GeneralizerParams,
		OmitHeader:        OmitHeader,
		OmitPadding:       OmitPadding,
		RightAlignNumeric: RightAlignNumeric,
		Spreadsheet:       copy(Spreadsheet),
		HeaderAlignment:   HeaderAlignment,
	}
}

func New(opts ...StringifyOption) *Stringify {
	s := Default()

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Stringify) Copy() *Stringify {
	return &Stringify{
		Style:             s.Style.Copy(),
		GeneralizerParams: s.GeneralizerParams.Copy(),
		OmitHeader:        s.OmitHeader,
		OmitPadding:       s.OmitPadding,
		RightAlignNumeric: s.RightAlignNumeric,
		Spreadsheet:       copy(s.Spreadsheet),
	}
}

func (s *Stringify) With(opts ...StringifyOption) *Stringify {
	ss := s.Copy()

	for _, opt := range opts {
		opt(ss)
	}

	return ss
}

type StringifyOption func(*Stringify)

func WithStyle(style *TableStyle) StringifyOption {
	return func(s *Stringify) {
		s.Style = style
	}
}

func WithGeneralizer(params *generalizer.Params) StringifyOption {
	return func(s *Stringify) {
		s.GeneralizerParams = params
	}
}

func WithOmitHeader(omit bool) StringifyOption {
	return func(s *Stringify) {
		s.OmitHeader = omit
	}
}

func WithOmitPadding(omit bool) StringifyOption {
	return func(s *Stringify) {
		s.OmitPadding = omit
	}
}

func WithRightAlignNumeric(right bool) StringifyOption {
	return func(s *Stringify) {
		s.RightAlignNumeric = right
	}
}

func WithSpreadsheet(header string) StringifyOption {
	return func(s *Stringify) {
		s.Spreadsheet = &header
	}
}

func WithHeaderAlignment(align alignment) StringifyOption {
	return func(s *Stringify) {
		s.HeaderAlignment = align
	}
}

func WithHeaderFormatter(f func(string) string) StringifyOption {
	return func(s *Stringify) {
		s.HeaderFormatter = f
	}
}

func WithAutoFormatHeader() StringifyOption {
	return func(s *Stringify) {
		s.HeaderFormatter = AutoFormatHeader
	}
}
