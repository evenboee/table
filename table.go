package table

import (
	"fmt"

	"github.com/evenboee/table/generalizer"
)

type Table string

func (t Table) String() string {
	return string(t)
}

func (t Table) Print() {
	fmt.Println(t.String())
}

func (s *Stringify) String(v any) Table {
	return s.Generate(generalizer.Any(v))
}

func String(v any, opts ...StringifyOption) Table {
	return New(opts...).String(v)
}
