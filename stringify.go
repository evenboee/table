package table

import (
	"strconv"
	"strings"
	"unicode"
)

// TODO: Account for left, right, and separators longer than 1 character e.g. "||" or ">>"

func Generate(headers []string, rows [][]string, opts ...ParamsOption) Table {
	return New(opts...).Generate(headers, rows)
}

func (s *Params) Generate(headers []string, rows [][]string) Table {
	validateTable(headers, rows)

	if s.HeaderFormatter != nil {
		for i, header := range headers {
			headers[i] = s.HeaderFormatter(header)
		}
	}

	if s.Spreadsheet != nil {
		headers = append([]string{*s.Spreadsheet}, headers...)
		for i, row := range rows {
			rows[i] = append([]string{strconv.Itoa(i + 1)}, row...)
		}

	}

	// find longest string in each column
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = unicodeLength(header)
	}
	for _, row := range rows {
		for i, cell := range row {
			if l := unicodeLength(cell); l > colWidths[i] {
				colWidths[i] = l
			}
		}
	}

	var sb strings.Builder

	// add top separator
	if s.Style.TopSeparator != nil {
		sb.WriteString(buildSeparator(colWidths, *s.Style.TopSeparator, s.Style.Padding, s.OmitPadding))
		sb.WriteString("\n")
	}

	if !s.OmitHeader {
		// add header
		sb.WriteString(s.Style.Left)
		for i, header := range headers {
			if i != 0 {
				sb.WriteString(s.Style.Separator)
			}
			sb.WriteString(cellString(header, colWidths[i], s.HeaderAlignment, s.Style.Padding, s.OmitPadding))
		}
		sb.WriteString(s.Style.Right)
		sb.WriteString("\n")

		// add header separator
		if s.Style.HeaderSeparator != nil {
			sb.WriteString(buildSeparator(colWidths, *s.Style.HeaderSeparator, s.Style.Padding, s.OmitPadding))
			sb.WriteString("\n")
		}
	}

	// add rows
	rowSep := ""
	if s.Style.RowSeparator != nil {
		rowSep = buildSeparator(colWidths, *s.Style.RowSeparator, s.Style.Padding, s.OmitPadding) + "\n"
	}

	for i, row := range rows {
		if i != 0 && rowSep != "" {
			sb.WriteString(rowSep)
		}

		sb.WriteString(s.Style.Left)
		for j, cell := range row {
			if j != 0 {
				sb.WriteString(s.Style.Separator)
			}

			align := Left
			if s.RightAlignNumeric && isNumeric(cell) {
				align = Right
			}

			sb.WriteString(cellString(cell, colWidths[j], align, s.Style.Padding, s.OmitPadding))
		}
		sb.WriteString(s.Style.Right)
		if i != len(rows)-1 {
			sb.WriteString("\n")
		}
	}

	// add bottom separator
	if s.Style.BottomSeparator != nil {
		sb.WriteString("\n")
		sb.WriteString(buildSeparator(colWidths, *s.Style.BottomSeparator, s.Style.Padding, s.OmitPadding))
	}

	return Table(sb.String())
}

func validateTable(headers []string, rows [][]string) {
	for _, row := range rows {
		if len(row) != len(headers) {
			panic("row length does not match header length")
		}
	}
}

func buildSeparator(colWidths []int, sep TableStyleSeparator, padding int, omitPadding bool) string {
	var sb strings.Builder
	sb.WriteString(sep.Left)
	for i, width := range colWidths {
		if i != 0 {
			sb.WriteString(sep.Junction)
		}
		if omitPadding {
			sb.WriteString(sep.Center)
		} else {
			sb.WriteString(strings.Repeat(sep.Center, width+padding*2))
		}
	}
	sb.WriteString(sep.Right)
	return sb.String()
}

func unicodeLength(s string) int {
	return len([]rune(s))
}

type alignment int

const (
	Left alignment = iota
	Right
	Center
)

func pad(s string, width int, align alignment) string {
	switch align {
	case Left:
		return s + strings.Repeat(" ", width-unicodeLength(s))
	case Right:
		return strings.Repeat(" ", width-unicodeLength(s)) + s
	case Center:
		diff := width - unicodeLength(s)
		left := diff / 2
		right := diff - left
		return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
	}
	return s
}

func cellString(s string, width int, align alignment, padding int, omitPadding bool) string {
	if omitPadding {
		return s
	}

	s = pad(s, width, align)
	p := strings.Repeat(" ", padding)

	return p + s + p
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// AutoFormatHeader takes a string and returns a string with the first letter of each word capitalized
// it also attemts to insert spaces between words that are not separated by spaces or underscores
// e.g. "hello_world" -> "Hello World"
// e.g. "HelloWorld" -> "Hello World"
// e.g. "hello-world" -> "Hello World"
func AutoFormatHeader(h string) string {
	hc := []rune(h)

	var result strings.Builder
	l := len(hc) - 1
	nextUpper := false
	for i, r := range hc {
		if r == '_' || r == '-' {
			result.WriteRune(' ')
			nextUpper = true
			continue
		}

		// if current letter is uppercase, not first,
		//   and next letter is lowercase
		//   or if letter is last and previous letter is lowercase:
		//   insert ' ' before current letter
		if unicode.IsUpper(r) && i != 0 &&
			(i < l && unicode.IsLower(rune(hc[i+1])) ||
				i == l && unicode.IsLower(rune(hc[i-1]))) {
			result.WriteRune(' ')
		}

		if nextUpper {
			r = unicode.ToUpper(r)
			nextUpper = false
		}
		result.WriteRune(r)
	}

	return result.String()
}
