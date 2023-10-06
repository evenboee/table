package table

type stringer struct {
	// Margin is the number of spaces between the data of a column and the edge or separators
	Margin int
	// LeftEdge is the left edge of the table
	LeftEdge string
	// RightEdge is the right edge of the table
	RightEdge string
	// Sep is the separator between columns
	Sep string
	// OmitHeaders determines whether the headers should be printed
	OmitHeaders bool
	// OmitPadding determines whether the padding should be printed
	OmitPadding bool
	// HeaderAlignment determines the alignment of the headers
	HeaderAlignment textAlignment
	// RowSep is the separator between rows
	RowSep *Separator
	// HeaderSep is the separator between the headers and the data
	HeaderSep *Separator
	// TopSep is the separator at the top of the table
	TopSep *Separator
	// BottomSep is the separator at the bottom of the table
	BottomSep *Separator

	// Should numeric values be right aligned
	RightAlignNumeric bool
}

func copySeparator(s *Separator) *Separator {
	if s == nil {
		return nil
	}
	return s.Copy()
}

func (s *stringer) Copy() *stringer {
	return &stringer{
		Margin:            s.Margin,
		LeftEdge:          s.LeftEdge,
		RightEdge:         s.RightEdge,
		Sep:               s.Sep,
		OmitHeaders:       s.OmitHeaders,
		OmitPadding:       s.OmitPadding,
		HeaderAlignment:   s.HeaderAlignment,
		RowSep:            copySeparator(s.RowSep),
		HeaderSep:         copySeparator(s.HeaderSep),
		TopSep:            copySeparator(s.TopSep),
		BottomSep:         copySeparator(s.BottomSep),
		RightAlignNumeric: s.RightAlignNumeric,
	}
}

var DefaultStringer = StringerCustom

// Custom layout
var StringerCustom = &stringer{
	Margin:    1,
	LeftEdge:  "|",
	RightEdge: "|",
	Sep:       "|",
	HeaderSep: &Separator{
		Left:     "+",
		Right:    "+",
		Main:     "-",
		Junction: "+",
	},
	RightAlignNumeric: true,
}

// Table layouts based on https://ozh.github.io/ascii-tables/

var stringerASCIIMySQLStyleHeaderSep = &Separator{
	Left:     "+",
	Right:    "+",
	Main:     "-",
	Junction: "+",
}

var StringerASCIIMySQLStyle = &stringer{
	Margin:            1,
	LeftEdge:          "|",
	RightEdge:         "|",
	Sep:               "|",
	RowSep:            nil,
	HeaderSep:         stringerASCIIMySQLStyleHeaderSep,
	TopSep:            stringerASCIIMySQLStyleHeaderSep,
	BottomSep:         stringerASCIIMySQLStyleHeaderSep,
	RightAlignNumeric: true,
}

var stringerASCIISeparatedRowSep = &Separator{
	Left:     "+",
	Right:    "+",
	Main:     "-",
	Junction: "+",
}

var stringerASCIISeparatedHeaderSep = &Separator{
	Left:     "+",
	Right:    "+",
	Main:     "=",
	Junction: "+",
}

var StringerASCIISeparated = &stringer{
	Margin:            1,
	LeftEdge:          "|",
	RightEdge:         "|",
	Sep:               "|",
	HeaderSep:         stringerASCIISeparatedHeaderSep,
	RowSep:            stringerASCIISeparatedRowSep,
	TopSep:            stringerASCIISeparatedHeaderSep,
	BottomSep:         stringerASCIISeparatedRowSep,
	RightAlignNumeric: true,
}

var StringerASCIICompact = &stringer{
	Margin:    1,
	LeftEdge:  " ",
	RightEdge: " ",
	Sep:       " ",
	HeaderSep: &Separator{
		Left:     " ",
		Right:    " ",
		Main:     "-",
		Junction: " ",
	},
	RightAlignNumeric: true,
}

var StringerGithubMarkdown = &stringer{
	Margin:    1,
	LeftEdge:  "|",
	RightEdge: "|",
	Sep:       "|",
	HeaderSep: &Separator{
		Left:     "|",
		Right:    "|",
		Main:     "-",
		Junction: "|",
	},
	RightAlignNumeric: true,
}

var StringerUnicode = &stringer{
	Margin:          2,
	HeaderAlignment: TextAlignmentCenter,
	LeftEdge:        "║",
	RightEdge:       "║",
	Sep:             "║",
	HeaderSep: &Separator{
		Left:     "╠",
		Right:    "╣",
		Main:     "═",
		Junction: "╬",
	},
	TopSep: &Separator{
		Left:     "╔",
		Right:    "╗",
		Main:     "═",
		Junction: "╦",
	},
	BottomSep: &Separator{
		Left:     "╚",
		Right:    "╝",
		Main:     "═",
		Junction: "╩",
	},
	RightAlignNumeric: true,
}

var StringerUnicodeSingleLine = &stringer{
	Margin:          2,
	HeaderAlignment: TextAlignmentCenter,
	LeftEdge:        "┃",
	RightEdge:       "┃",
	Sep:             "┃",
	HeaderSep: &Separator{
		Left:     "┣",
		Right:    "┫",
		Main:     "━",
		Junction: "╋",
	},
	TopSep: &Separator{
		Left:     "┏",
		Right:    "┓",
		Main:     "━",
		Junction: "┳",
	},
	BottomSep: &Separator{
		Left:     "┗",
		Right:    "┛",
		Main:     "━",
		Junction: "┻",
	},
	RightAlignNumeric: true,
}

var StringerHTML = &stringer{
	Margin:      0,
	OmitPadding: true,
	LeftEdge:    "<tr><td>",
	RightEdge:   "</td></tr>",
	Sep:         "</td><td>",
	HeaderSep: &Separator{
		Left:     "</thead>",
		Right:    "<tbody>",
		Main:     "",
		Junction: "",
	},
	TopSep: &Separator{
		Left:     "<table>",
		Right:    "<thead>",
		Main:     "",
		Junction: "",
	},
	BottomSep: &Separator{
		Left:     "</tbody>",
		Right:    "</table>",
		Main:     "",
		Junction: "",
	},
	RightAlignNumeric: false,
}
