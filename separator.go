package table

import "strings"

func buildSeparator(sep *Separator, colWidths []int) string {
	if sep == nil {
		return ""
	}
	return sep.toString(colWidths)
}

type Separator struct {
	Left     string
	Right    string
	Junction string
	Main     string
}

func (s *Separator) toString(colWidths []int) string {
	var sb strings.Builder

	sb.WriteString(s.Left)

	for i, width := range colWidths {
		if i != 0 {
			sb.WriteString(s.Junction)
		}

		sb.WriteString(strings.Repeat(s.Main, width))
	}

	sb.WriteString(s.Right)
	sb.WriteByte('\n')

	return sb.String()
}

func (s *Separator) Copy() *Separator {
	return &Separator{
		Left:     s.Left,
		Right:    s.Right,
		Junction: s.Junction,
		Main:     s.Main,
	}
}
