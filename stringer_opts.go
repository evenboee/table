package table

type stringerOpt func(*stringer)

func WithMargin(margin int) stringerOpt {
	return func(s *stringer) {
		s.Margin = margin
	}
}

func WithHeaderSep(headerSep *Separator) stringerOpt {
	return func(s *stringer) {
		s.HeaderSep = headerSep
	}
}

func WithRowSep(rowSep *Separator) stringerOpt {
	return func(s *stringer) {
		s.RowSep = rowSep
	}
}

func WithTopSep(topSep *Separator) stringerOpt {
	return func(s *stringer) {
		s.TopSep = topSep
	}
}

func WithBottomSep(bottomSep *Separator) stringerOpt {
	return func(s *stringer) {
		s.BottomSep = bottomSep
	}
}

func WithRightAlignNumeric(rightAlignNumeric bool) stringerOpt {
	return func(s *stringer) {
		s.RightAlignNumeric = rightAlignNumeric
	}
}

func WithOmitHeaders(omitHeaders bool) stringerOpt {
	return func(s *stringer) {
		s.OmitHeaders = omitHeaders
	}
}

func WithOmitPadding(omitPadding bool) stringerOpt {
	return func(s *stringer) {
		s.OmitPadding = omitPadding
	}
}

func WithHeaderAlignment(headerAlignment textAlignment) stringerOpt {
	return func(s *stringer) {
		s.HeaderAlignment = headerAlignment
	}
}

func WithCenteredHeaders(s *stringer) {
	s.HeaderAlignment = TextAlignmentCenter
}

func WithLeftEdge(leftEdge string) stringerOpt {
	return func(s *stringer) {
		s.LeftEdge = leftEdge
	}
}

func WithRightEdge(rightEdge string) stringerOpt {
	return func(s *stringer) {
		s.RightEdge = rightEdge
	}
}

func WithSep(sep string) stringerOpt {
	return func(s *stringer) {
		s.Sep = sep
	}
}

type sep uint8

const (
	SepHeader sep = iota
	SepRow
	SepTop
	SepBottom
)

// WithSepAsSep sets the separator of dst to the separator of src
func WithSepAsSep(dst, src sep) stringerOpt {
	return func(str *stringer) {
		var s *Separator
		switch src {
		case SepHeader:
			s = str.HeaderSep
		case SepRow:
			s = str.RowSep
		case SepTop:
			s = str.TopSep
		case SepBottom:
			s = str.BottomSep
		}

		switch dst {
		case SepHeader:
			str.HeaderSep = s
		case SepRow:
			str.RowSep = s
		case SepTop:
			str.TopSep = s
		case SepBottom:
			str.BottomSep = s
		}
	}
}

func (s *stringer) WithOpts(opts ...stringerOpt) *stringer {
	stringer := s.Copy()

	for _, opt := range opts {
		opt(stringer)
	}

	return stringer
}

func NewStringer(opts ...stringerOpt) *stringer {
	s := &stringer{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
