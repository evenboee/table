package table

type TableStyle struct {
	// Padding is the minimum number of spaces left and right of the content within a cell
	Padding int

	// Left is the left side of the row
	Left string
	// Right is the right side of the row
	Right string
	// Separator is the separator between columns
	Separator string

	// TopSeparator is the separator at the top of the table
	TopSeparator *TableStyleSeparator
	// HeaderSeparator is the separator between the header and the body
	HeaderSeparator *TableStyleSeparator
	// BottomSeparator is the separator at the bottom of the table
	BottomSeparator *TableStyleSeparator
	// RowSeparator is the separator between rows
	RowSeparator *TableStyleSeparator
}

func (s *TableStyle) Copy() *TableStyle {
	return &TableStyle{
		Padding:         s.Padding,
		Left:            s.Left,
		Right:           s.Right,
		Separator:       s.Separator,
		TopSeparator:    s.TopSeparator.Copy(),
		HeaderSeparator: s.HeaderSeparator.Copy(),
		BottomSeparator: s.BottomSeparator.Copy(),
		RowSeparator:    s.RowSeparator.Copy(),
	}
}

func (s *TableStyle) With(f func(*TableStyle)) *TableStyle {
	sc := s.Copy()
	f(sc)
	return sc
}

type TableStyleSeparator struct {
	Left     string // "<"
	Right    string // ">"
	Junction string // "+"
	Center   string // "-"

	// <---+---+--->
}

func (s *TableStyleSeparator) Copy() *TableStyleSeparator {
	if s == nil {
		return nil
	}

	return &TableStyleSeparator{
		Left:     s.Left,
		Right:    s.Right,
		Junction: s.Junction,
		Center:   s.Center,
	}
}

// Table styles based on https://ozh.github.io/ascii-tables/

var DefaultStyle = StyleASCIIMySQLStyle

var styleASCIIMySQLStyleHeaderSep = &TableStyleSeparator{
	Left:     "+",
	Right:    "+",
	Junction: "+",
	Center:   "-",
}

var StyleASCIIMySQLStyle = &TableStyle{
	Padding:         1,
	Left:            "|",
	Right:           "|",
	Separator:       "|",
	RowSeparator:    nil,
	HeaderSeparator: styleASCIIMySQLStyleHeaderSep,
	TopSeparator:    styleASCIIMySQLStyleHeaderSep,
	BottomSeparator: styleASCIIMySQLStyleHeaderSep,
}

var styleASCIISeparatedRowSep = &TableStyleSeparator{
	Left:     "+",
	Right:    "+",
	Junction: "+",
	Center:   "-",
}

var styleASCIISeparatedHeaderSep = &TableStyleSeparator{
	Left:     "+",
	Right:    "+",
	Junction: "+",
	Center:   "=",
}

var StyleASCIISeparated = &TableStyle{
	Padding:         1,
	Left:            "|",
	Right:           "|",
	Separator:       "|",
	TopSeparator:    styleASCIISeparatedHeaderSep,
	HeaderSeparator: styleASCIISeparatedHeaderSep,
	RowSeparator:    styleASCIISeparatedRowSep,
	BottomSeparator: styleASCIISeparatedRowSep,
}

var StyleASCIICompact = &TableStyle{
	Padding:   1,
	Left:      " ",
	Right:     " ",
	Separator: " ",
	HeaderSeparator: &TableStyleSeparator{
		Left:     " ",
		Right:    " ",
		Junction: " ",
		Center:   "-",
	},
}

var StyleGithubMarkdown = &TableStyle{
	Padding:   1,
	Left:      "|",
	Right:     "|",
	Separator: "|",
	HeaderSeparator: &TableStyleSeparator{
		Left:     "|",
		Right:    "|",
		Junction: "|",
		Center:   "-",
	},
}

var StyleUnicode = &TableStyle{
	Padding:   2,
	Left:      "║",
	Right:     "║",
	Separator: "║",
	HeaderSeparator: &TableStyleSeparator{
		Left:     "╠",
		Right:    "╣",
		Junction: "╬",
		Center:   "═",
	},
	TopSeparator: &TableStyleSeparator{
		Left:     "╔",
		Right:    "╗",
		Junction: "╦",
		Center:   "═",
	},
	BottomSeparator: &TableStyleSeparator{
		Left:     "╚",
		Right:    "╝",
		Junction: "╩",
		Center:   "═",
	},
}

var StyleUnicodeSingleLine = &TableStyle{
	Padding:   2,
	Left:      "┃",
	Right:     "┃",
	Separator: "┃",
	HeaderSeparator: &TableStyleSeparator{
		Left:     "┣",
		Right:    "┫",
		Junction: "╋",
		Center:   "━",
	},
	TopSeparator: &TableStyleSeparator{
		Left:     "┏",
		Right:    "┓",
		Junction: "┳",
		Center:   "━",
	},
	BottomSeparator: &TableStyleSeparator{
		Left:     "┗",
		Right:    "┛",
		Junction: "┻",
		Center:   "━",
	},
}

var StyleHTML = &TableStyle{
	Padding:   0,
	Left:      "<tr><td>",
	Right:     "</td></tr>",
	Separator: "</td><td>",
	HeaderSeparator: &TableStyleSeparator{
		Left:     "</thead>",
		Right:    "<tbody>",
		Junction: "",
		Center:   "",
	},
	TopSeparator: &TableStyleSeparator{
		Left:     "<table>",
		Right:    "<thead>",
		Junction: "",
		Center:   "",
	},
	BottomSeparator: &TableStyleSeparator{
		Left:     "</tbody>",
		Right:    "</table>",
		Junction: "",
		Center:   "",
	},
}
