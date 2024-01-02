package table

import (
	"github.com/evenboee/table/generalizer"
	"strconv"
	"strings"
)

func (s *StringifyConfig) Stringify(res generalizer.Result) string {
	headers := res.Headers
	data := res.Rows

	if s.Spreadsheet != nil {
		headers, data = insertSpreadsheetColumn(*s.Spreadsheet, headers, data)
	}

	longestOfColumn := longestValueOfColumn(headers, data)

	rowLength := unicodeLength(s.LeftEdge) + unicodeLength(s.RightEdge)
	additionalLength := s.Margin*2 + unicodeLength(s.Sep)
	for _, v := range longestOfColumn {
		rowLength += v + additionalLength
	}

	margin := strings.Repeat(" ", s.Margin)
	sep := margin + s.Sep + margin
	leadingEdge := s.LeftEdge + margin
	trailingEdge := margin + s.RightEdge

	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = longestOfColumn[header] + s.Margin*2
	}

	// Building separators
	headerSepRow := buildSeparator(s.HeaderSep, colWidths) // getSeparatorRow(longestOfColumn, headers, s.HeaderSepJunction, s.HeaderSep, len(s.Edge), len(s.Sep), s.Margin)
	rowSep := buildSeparator(s.RowSep, colWidths)
	topSep := buildSeparator(s.TopSep, colWidths)
	bottomSep := buildSeparator(s.BottomSep, colWidths)

	var sb strings.Builder

	// Writing top separator
	sb.WriteString(topSep)

	if !s.OmitHeaders {
		// Writing header row
		sb.WriteString(leadingEdge)
		for i, header := range headers {
			if i != 0 {
				sb.WriteString(sep)
			}

			if s.OmitPadding {
				sb.WriteString(header)
			} else {
				sb.WriteString(getAlignedString(header, longestOfColumn[header], s.HeaderAlignment))
			}
		}
		sb.WriteString(trailingEdge)
		sb.WriteByte('\n')

		// Writing header separator row
		sb.WriteString(headerSepRow)
	}

	// Writing data rows
	lenData := len(data)
	maxIdx := lenData
	if s.Limit > 0 && s.Offset+s.Limit < lenData {
		maxIdx = s.Offset + s.Limit
	}
	for i := s.Offset; i < maxIdx; i++ {
		row := data[i]
		sb.WriteString(leadingEdge)
		for j, header := range headers {
			if j != 0 {
				sb.WriteString(sep)
			}

			content := row[header]
			if s.OmitPadding {
				sb.WriteString(content)
			} else {
				alignment := TextAlignmentLeft
				if s.RightAlignNumeric && isNumeric(content) {
					alignment = TextAlignmentRight
				}

				sb.WriteString(getAlignedString(content, longestOfColumn[header], alignment))
			}
		}
		sb.WriteString(trailingEdge)
		sb.WriteByte('\n')
		if i != len(data)-1 {
			sb.WriteString(rowSep)
		}
	}

	// Writing bottom separator row
	sb.WriteString(bottomSep)

	return sb.String()
}

func insertSpreadsheetColumn(columnHeader string, headers []string, data []map[string]string) ([]string, []map[string]string) {
	headers = append([]string{columnHeader}, headers...)
	for i := range data {
		data[i][columnHeader] = strconv.Itoa(i + 1)
	}

	return headers, data
}

type textAlignment uint8

const (
	TextAlignmentLeft textAlignment = iota // Default
	TextAlignmentRight
	TextAlignmentCenter
)

func getAlignedString(text string, width int, alignment textAlignment) string {
	l := unicodeLength(text)
	if l >= width {
		return text
	}

	switch alignment {
	case TextAlignmentLeft:
		return text + getFillerString(l, width)
	case TextAlignmentRight:
		return getFillerString(l, width) + text
	case TextAlignmentCenter:
		leftoverWidth := width - l
		left := leftoverWidth / 2
		right := leftoverWidth - left
		return getFillerString(left, leftoverWidth) + text + getFillerString(right, leftoverWidth)
	}

	return text
}

func getFillerString(actual int, expected int) string {
	if actual >= expected {
		return ""
	}

	return strings.Repeat(" ", expected-actual)
}

func longestValueOfColumn(headers []string, data []map[string]string) map[string]int {
	longestOfHeader := make(map[string]int)
	for _, header := range headers {
		longestOfHeader[header] = unicodeLength(header)
	}

	for _, row := range data {
		for k, v := range row {
			if l := unicodeLength(v); l > longestOfHeader[k] {
				longestOfHeader[k] = l
			}
		}
	}

	return longestOfHeader
}

func unicodeLength(s string) int {
	return len([]rune(s))
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
